package vo

import (
	"gomo/db/models"
)

type SubjectVO struct {
	CatID   int64
	CatTitle string
	Items   []models.Item
}
