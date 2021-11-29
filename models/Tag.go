package models

import "time"

type Tag struct {
	ID int64
	Name string
	CreateTime time.Time
	UpdateTime time.Time
}
