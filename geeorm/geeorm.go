package geeorm

import (
	"GReORM/session"
	"database/sql"
	"log"
)

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
	//Sending a ping to make sure that database connection is alive
	if err = db.Ping(); err != nil {
		log.Println(err)
		return
	}
	e = &Engine{db: db}
	log.Println("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Println("Failed to close database")
	}
	log.Println("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Panic(err)
		return
	}
	//Send a ping to make sure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Panic(err)
		return
	}
	//make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Panicf("dialect %s not found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Println("Connect database success")
	return
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
