package vo

import "leetroll/db/models"

type ChapterVO struct {
	ID       int64
	Chapter  int64
	Main     []models.File
	Episodes []models.File
}
