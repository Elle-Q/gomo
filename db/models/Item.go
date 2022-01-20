package models

import "time"

type Item struct {
	ID int64
	Name string
	Desc string
	Preview string
	Type string
	BLink string
	Tags string
	Price int64
	Author int64
	Class string
	DownCnt int64
	Scores int64
	CreateTime time.Time
	UpdateTime time.Time
}