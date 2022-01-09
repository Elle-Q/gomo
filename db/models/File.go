package models

import (
	"database/sql"
	"time"
)

type File struct {
	ID int64
	ItemId sql.NullInt64
	Name string
	Type string
	QnLink string
	Size float32
	Format string
	CreateTime time.Time
	UpdateTime time.Time
}
