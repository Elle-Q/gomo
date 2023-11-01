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

type ItemWithFilesVO struct {
	ID         int64
	Name       string
	Type       string
	Desc       string
	Tags       []string
	Author     string
	Scores     float64
	Price      float64
	CatID      int64
	CatTitle   string
	DownCnt    int64
	Main       string
	Attachment []models.File
	Preview    []models.File
	Chapters   []ChapterVO
}
