package models

import (
	// r "github.com/dancannon/gorethink"
	// re "github.com/dancannon/gorethink/encoding"
	"time"
)

type User struct {
	UserID    string
	NickName  string
	Hash      string
	CreatedAt time.Time
}
