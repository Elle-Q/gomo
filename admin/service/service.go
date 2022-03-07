package service

import "gomo/db/handlers"

func NewItemService() ItemService{
	return ItemService{
		ItemHandler: &handlers.ItemHandler{},
		FileHandler: &handlers.FileHandler{},
	}
}

