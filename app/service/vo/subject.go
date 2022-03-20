package vo

import (
	"gomo/db/models"
)

type SubjectVO struct {
	CatID    int64
	CatTitle string
	CatSubTitle  string
	Items    []models.Item
}
