package models

import (
	"time"
)

type File struct {
	ID         int64
	ItemId     int64 //sql.NullInt64
	Name       string
	Type       string
	QnLink     string
	Size       float32
	Format     string
	Bucket     string
	Key        string
	CreateTime time.Time
	UpdateTime time.Time
}
