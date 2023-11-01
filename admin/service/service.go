package service

import "leetroll/db/handlers"

func NewItemService() ItemService {
	return ItemService{
		ItemHandler:    &handlers.ItemHandler{},
		FileHandler:    &handlers.FileHandler{},
		ChapterHandler: &handlers.ChapterHandler{},
	}
}
