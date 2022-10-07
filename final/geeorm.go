package main

import (
	"database/sql"
	"final/log"
	"final/session"
)

/*
  Session负责与数据库的交互
  交互前的准备工作和交互后的收尾工作（关闭连接）则由Session负责。
  Engine是GeeORM与用户交互的入口
*/

type Engine struct {
	db *sql.DB
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
	//new engine
	e = &Engine{db: db}
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
	return session.New(engine.db)
}
