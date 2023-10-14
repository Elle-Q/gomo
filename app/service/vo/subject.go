package vo

import (
	"leetroll/db/models"
)

type SubjectVO struct {
	CatID       int64
	CatTitle    string
	CatSubTitle string
	Items       []models.Item
}
