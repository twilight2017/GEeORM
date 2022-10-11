package session

import (
	"final/clause"
	"reflect"
)

//support map[string]interface{}
//also support kv list: "Name", "Tom", "Age", 18
func (s *Session) Update(kv ...interface{}) (int64, error) {
	m, ok := kv[0].(map[string]interface{}) //Update方法会动态地判断传入参数的类型，如果不是map类型，会进行自动转换
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv(i).(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.RefTable().Name(), m)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.RAW(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

//Delete records with where clause
func (s *Session) Delete() (int64, error) {
	s.clause.Set(clause.DELETE, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	result, err := s.RAW(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

//Count records with where clause
func (s Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, vars := sq.clause.Build(clause.COUNT, clause.WHERE)
	result, err := s.RAW(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

//Limit adds limit condition to clause
func (s *Session) Limit(num int) *Session {
	s.clause.Set(clause.LIMIT, num)
	return s
}

//Where adds limit condition to clause
func (s *Session) Where(desc string, args ...interface{}) *Session {
	var vars []interface{}
	s.clause.Set(clause.WHERE, append(append(vars, desc), args...)...)
	return s
}

//OrderBy adds order by condition to clause
func (s *Session) OrderBy(desc string) *Session {
	s.clause.Set(clause.ORDERBY, desc)
	return s
}

func (s *Session) First(value interfcae{}) error{
	dest := reflect.Indirect(reflect.Valueof(value))
	destSlice := reflect.New(reflect.Sliceof(dest.Type())).Elem()
	if err := s.Limit(1).Find(destSlice.Addr().Interface()); err != nil{
		return err
	}
	if destSlice.Len() == 0{
		return errors.New("NOT FOUND")
	}
	dest.Set(destSlice.Index(0))
	return nil
}