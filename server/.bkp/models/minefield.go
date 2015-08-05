package models

import (
	r "github.com/dancannon/gorethink"
	re "github.com/dancannon/gorethink/encoding"
	"time"
)

type MineField struct {
	MatrixID  string    `gorethink:"id,omitempty"`
	Matrix    [][]int   `gorethink:"matrix"`
	CreatedAt time.Time `gorethink:"createdAt"`
	CreatedBy string    `gorethink:"createdBy"`
}

func (e *Engine) SaveMineField(mat [][]int, userID string) (*MineField, error) {
	mf := MineField{
		Matrix:    mat,
		CreatedAt: time.Now().UTC(),
		CreatedBy: userID,
	}

	insert, insertErr := e.term.Table(MineTB).
		Insert(&mf, r.InsertOpts{ReturnChanges: true, Durability: "hard"}).
		RunWrite(e.session)
	if insertErr != nil {
		e.log.Println(insertErr)
		return nil, ErrQueryFailed
	}
	if insert.Inserted != 1 {
		e.log.Println(insert.FirstError)
		return nil, ErrQueryFailed
	}
	if scanErr := re.Decode(&mf, insert.Changes[0].NewValue); scanErr != nil {
		e.log.Println(scanErr)
		return nil, ErrQueryFailed
	}

	return &mf, nil
}

func (e *Engine) GetMineFields() ([]MineField, error) {
	get, getErr := e.term.Table(MineTB).
		Run(e.session)
	if getErr != nil {
		e.log.Println(getErr)
		return nil, ErrQueryFailed
	}

	fields := make([]MineField, 0)
	if scanErr := get.All(&fields); scanErr != nil {
		e.log.Println(scanErr)
		return nil, ErrQueryFailed
	}

	return fields, nil
}

func (e *Engine) GetMineFieldByID(fieldID string) (*MineField, error) {
	if len(fieldID) == 0 {
		return nil, ErrInvalidFieldID
	}

	get, getErr := e.term.Table(MineTB).
		Get(fieldID).
		Run(e.session)
	if getErr != nil {
		e.log.Println(getErr)
		return nil, ErrQueryFailed
	}

	var field MineField
	if scanErr := get.One(&field); scanErr != nil {
		e.log.Println(scanErr)
		return nil, ErrQueryFailed
	}

	return &field, nil
}

func (e *Engine) GetMineFieldsByUserID(userID string) ([]MineField, error) {
	if len(userID) == 0 {
		return nil, ErrInvalidFieldID
	}

	get, getErr := e.term.Table(MineTB).
		GetAllByIndex(FieldCreatedBy, userID).
		Run(e.session)
	if getErr != nil {
		e.log.Println(getErr)
		return nil, ErrQueryFailed
	}

	fields := make([]MineField, 0)
	if scanErr := get.All(&fields); scanErr != nil {
		e.log.Println(scanErr)
		return nil, ErrQueryFailed
	}

	return fields, nil
}
