package models

import "time"

type Config struct {
	ID int64
	Name string
	Val string
	CreateTime time.Time
	UpdateTime time.Time
}