package service

import (
	"gomo/db/handlers"
)

func NewCatItemService() CatItemService{
	return CatItemService{
		ItemHandler: &handlers.ItemHandler{},
		CatHandler: &handlers.CatHandler{},
	}
}

