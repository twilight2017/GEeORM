// Package session 在这一部分实现调用SQL语句实现原生交互
package session

import (
	"database/sql"
	"final/dialect"
	"final/log"
	"final/schema"
	"strings"
)

type Session struct {
	db *sql.DB //使用sql.Open()方法连接数据库成功后返回的指针
	//用于拼接sql语句和sql语句中占位符的对应值
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db: db, dialect: dialect}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

//封装Exec,QueryRow, QueryRows方法，加上log，和对Session的关闭操作
//封装有2个目的：1.统一打印日志（包括执行的SQL语句和错误日志） 2.清空(s *Session).sql和(s *Session).sqlVars两个变量，这样Session可以复用，开启一次绘画，可以执行多次SQL
//Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return result, err
}

//QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

//QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
