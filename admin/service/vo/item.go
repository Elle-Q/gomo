package vo

import "time"

type ItemVO struct {
	ID int64
	CatName int64
	Name string
	Desc string
	Preview string
	BLink string
	Tags string
	Price int64
	Author string
	DownCnt int64
	Scores int64
	CreateTime time.Time
	UpdateTime time.Time
}
