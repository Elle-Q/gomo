package service

import (
	"leetroll/db/handlers"
)

func NewCatItemService() CatItemService {
	return CatItemService{
		ItemHandler: &handlers.ItemHandler{},
		CatHandler:  &handlers.CatHandler{},
	}
}

func NewItemService() ItemService {
	return ItemService{
		ItemHandler: &handlers.ItemHandler{},
		FileHandler: &handlers.FileHandler{},
	}
}
