package models

import "time"

type SearchLog struct {
	ID int64
	UserId int64
	RawQuery string
	Words string
	CatId int64
	CreateTime time.Time
	UpdateTime time.Time
}