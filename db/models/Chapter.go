package models

import "time"

type Chapter struct {
	ID         int64
	ItemId     int64
	Main       int64
	Chapter    int64
	CreateTime time.Time
	UpdateTime time.Time
}
