package vo

import "leetroll/db/models"

type ChapterVO struct {
	ID       int64
	Chapter  int64
	Main     string
	Episodes []models.File
}
