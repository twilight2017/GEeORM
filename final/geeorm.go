package main

import (
	"database/sql"
	"final/dialect"
	"final/log"
	"final/session"
)

/*
  Session负责与数据库的交互
  交互前的准备工作和交互后的收尾工作（关闭连接）则由Session负责。
  Engine是GeeORM与用户交互的入口
*/

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	//sending a ping to make sure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	//make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

//Close the database
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to cloe database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
