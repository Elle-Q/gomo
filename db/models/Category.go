package models

import "time"

type Category struct {
	ID int64
	Title string
	SubTitle string
	Preview string
	Desc string
	Status string
	CreateTime time.Time
	UpdateTime time.Time
}