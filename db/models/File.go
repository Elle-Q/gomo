package models

import "time"

type File struct {
	ID int64
	ItemId int64
	Name string
	QnLink string
	size float32
	format string
	CreateTime time.Time
	UpdateTime time.Time
}
