package models

import "time"

type Item struct {
	ID int64
	Cat *Category
	Name string
	Desc string
	Preview string
	Type string
	BLink string
	Tags string
	Price int64
	Author string
	DownCnt int64
	Scores int64
	Status string
	CreateTime time.Time
	UpdateTime time.Time
}