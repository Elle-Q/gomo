package models

import "time"

type Piece struct {
	ID int64
	title string
	desc string
	status string
	show bool
	likes int64
	hates int64
	CreateTime time.Time
	UpdateTime time.Time
}