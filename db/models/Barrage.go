package models

import "time"

type Barrage struct {
	ID int64
	ItemId int64
	SeekTime int64
	content string
	Show bool
	CreateTime time.Time
	UpdateTime time.Time
}
