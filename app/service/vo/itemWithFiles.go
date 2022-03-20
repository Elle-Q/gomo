package vo

import (
	"gomo/db/models"
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
	ID       int64
	Name     string
	RescType string
	Desc     string
	Tags     []string
	Author   string
	Scores   int64
	CatID    int64
	CatTitle string
	DownCnt  int64
	Main     []models.File
	Refs     []models.File
	Preview  []models.File
}
