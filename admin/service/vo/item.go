package vo

import (
	"leetroll/db/models"
	"time"
)

type ItemVO struct {
	ID         int64
	CatName    int64
	Name       string
	Desc       string
	Preview    string
	BLink      string
	Tags       string
	Price      int64
	Author     string
	DownCnt    int64
	Scores     int64
	CreateTime time.Time
	UpdateTime time.Time
}

type ItemFilesVO struct {
	ID         int64
	ItemName   string
	RescType   string
	Main       []models.File
	Attachment []models.File
	Preview    []models.File
	Chapters   []ChapterVO
}
