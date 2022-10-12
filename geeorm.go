package main

import (
	"final/log"
	"final/session"
)

type TxFunc func(s session.Session) (interface{}, error)

func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			err = s.Commit()
		}
	}()

	return f(s)
}

//difference returns a-b
func difference(a []string, b []string) (diff []string) {
	mapB := make(map[string]bool)
	for _, v := range b {
		mapB[v] = True
	}
	for _, v := range a {
		if _, ok := mapB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}

//Migrate table
func (engine *Engine) Migrate(value interface{}) error {
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		if !s.Model(value).HasTable() {
			log.Infof("table %s doesn't exist", s.RefTable().Name)
			return nil, s.CreateTable()
		}
		table := s.RefTable()
		rows, _ := s.RAW(fmt.Spintf("SELECT * FROM %s LIMIT 1", table.Name)).QueryRows()
		columns, _ := rows.Columns()
		//新表-旧表=新增字段
		addCols := difference(table.FieldNames, columns)
		//旧表-新表=删除字段
		delCols := difference(columns, table.FieldNames)
		log.Infof("added cols %v, deleted cols %v", addCols, delCols)

		for _, col := range addCols {
			f := table.GetField(col)
			sqlStr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", table.Name, f.Name, f.Type)
			if _, err = s.RAW(sqlStr).Exec(); err != nil {
				return
			}
		}
		if len(delCols) == 0 {
			return
		}
		tmp := "tmp_" + table.Name
		fieldStr := strings.Join(table.FieldNames, ", ")
		//从old table中选择要保留的字段到tmp_table中
		s.RAW(fmt.Sprintf("CREATE TABLE %s AS SELECT %s from %s", tmp, fieldStr, table.Name))
		s.RAW(fmt.Sprintf("DROP TABLE %s", table.Name))
		s.RAW(fmt.Sprintf("ALTER TABLE %s RENAME TO %s", tmp, table.Name))
		_, err = s.Exec()
		return
	})
	return error
}
