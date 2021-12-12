package models

import "time"

type User struct {
	ID int64
	Name string
	Phone string
	QRCode string
	Address string
	Gender string
	Vip bool
	BgImag string
	Admin bool
	Status string
	CreateTime time.Time
	UpdateTime time.Time

}