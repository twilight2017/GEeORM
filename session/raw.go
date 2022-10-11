//该部分实现直接使用SQL语句进行原生交互
package session

import (
	"GReORM/dialect"
	"GReORM/schema"
	"database/sql"
	"log"
	"strings"
)

type Session struct {
	db       *sql.DB //连接数据库后返回的指针
	dialect  dialect.Dialect
	tx       *sql.Tx //增加对事务的支持
	refTable *schema.Schema
	clause   clause.Clause
	//拼接sql语句和sql语句中占位符的对应值
	sql     strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

//get db
func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

//Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear() //清空sql和sqlVars两个变量，这样session可以复用，执行多次sql
	log.Println(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Println("Error Execution")
	}
	return
}

//QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Println(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

//QueryRows gets a list of records from db
func (s *Session) QuerRows() (rows *sql.rows, err error) {
	defer s.Clear()
	log.Println(s.sql.String(), s.sqlVars)
	if rows, err = s.DB.Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Println("error")
	}
	return
}

//CommonDB is a minimal function set for db
type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)
