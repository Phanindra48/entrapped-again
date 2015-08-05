package models

import (
	"errors"
)

var (
	ErrInvalidDB      = errors.New("invalid db name")
	ErrInvalidAddress = errors.New("invalid db connection address")
	ErrRethinkConn    = errors.New("unable to connect to rethink")
	ErrQueryFailed    = errors.New("unable to query database")
	ErrInvalidFieldID = errors.New("invalid minefield id")
)
