package models

import "time"

type Coment struct {
	ID int64
	UserId int64
	ItemId int64
	Content string
	Likes int64
	Hates int64
	IsParent bool
	Show bool
	CreateTime time.Time
	UpdateTime time.Time
}