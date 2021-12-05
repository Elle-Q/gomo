package models

import "time"

type Order struct {
	ID int64
	ItemId int64
	UserId int64
	status string
	price float32
	Type string
	CreateTime time.Time
	UpdateTime time.Time
}
