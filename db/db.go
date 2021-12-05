package db

import (
	"database/sql"
	"fmt"
)

type Handler struct {
	DB   *sql.DB
	Error error
}


func (db *Handler) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}