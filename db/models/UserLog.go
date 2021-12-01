package models

import "time"

type UserLog struct {
	ID int64
	UserId string
	CatId int64
	ItemId int64
	action string
	device string
	ip string
	info string
	CreateTime time.Time
	UpdateTime time.Time
}