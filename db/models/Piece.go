package models

import "time"

type Piece struct {
	ID int64
	Title string
	Desc string
	Status string
	Show bool
	Likes int64
	Hates int64
	CreateTime time.Time
	UpdateTime time.Time
}