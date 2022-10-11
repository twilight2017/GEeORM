package session

import "geeorm/log"

func (s *Session) Begin(err error) {
	log.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Commit(err error) {
	log.Info("transaction Commit")
	if err := s.tx.Commit(); err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction Rollback")
	if err := s.tx.Rollback(); err != nil {
		log.Error(err)
		return
	}
	return
}
