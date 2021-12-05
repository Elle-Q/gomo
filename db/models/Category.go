package models

import "time"

type Category struct {
	ID int64
	title string
	SubTitle string
	Preview string
	Desc string
	CreateTime time.Time
	UpdateTime time.Time
}