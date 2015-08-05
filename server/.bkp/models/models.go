package models

import (
	r "github.com/dancannon/gorethink"
	"log"
	"os"
	"time"
)

const (
	UserTB string = "users"
	MineTB string = "minefields"

	FieldCreatedBy string = "createdBy"
)

type Conf struct {
	DBAddress  string
	DBDatabase string
	DBAuthKey  string
	DBMaxIdle  int
	DBMaxOpen  int
	DBTimeout  time.Duration
}

type Engine struct {
	log     *log.Logger
	term    r.Term
	session *r.Session
}

func New(conf Conf) (*Engine, error) {
	if len(conf.DBAddress) == 0 {
		return nil, ErrInvalidAddress
	}
	if len(conf.DBDatabase) == 0 {
		return nil, ErrInvalidDB
	}

	logger := log.New(os.Stdout, "[entrapped models]", log.Ldate|log.Ltime|log.Lshortfile)

	session, sessionErr := r.Connect(r.ConnectOpts{
		Address:  conf.DBAddress,
		Database: conf.DBDatabase,
		AuthKey:  conf.DBAuthKey,
		Timeout:  conf.DBTimeout,
		MaxIdle:  conf.DBMaxIdle,
		MaxOpen:  conf.DBMaxOpen,
	})

	if sessionErr != nil {
		logger.Println(sessionErr)
		return nil, ErrRethinkConn
	}

	return &Engine{logger, r.DB(conf.DBDatabase), session}, nil
}
