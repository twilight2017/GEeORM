package session

import (
	"final/clause"
	"reflect"
)

//interface{}可以表示任意类型
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).RefTable()
		c.clause.Set(clause.INSERT, table.Name, table.FieldNames) //clause.Set()构造一个子句
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES) //clause.Build()按照传入的顺序构造出最终的SQL语句
	result, err := s.RAW(sql, vars...).Exec()//调用RAW().Exec()方法执行sql语句
	if err != nil {
		return 0, err
	}
	return result.RowAffected()
}

func (s *Session) Find(values interface{}) error{
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type.Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars = s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.RAW(sql, vars...).QueryRows()
	if err != nil{
		return err
	}
	for rows.Next(){
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for_, name := range table.FieldNames{
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := row.Scan(values...); err != nil{
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
