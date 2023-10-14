package service

import (
	"leetroll/app/service/vo"
	"leetroll/db/handlers"
	"leetroll/db/models"
)

type CatItemService struct {
	ItemHandler *handlers.ItemHandler
	CatHandler  *handlers.CatHandler
	Error       error
}

func (c *CatItemService) ListCatsWithItems(vos *[]vo.SubjectVO) *CatItemService {
	itemHandler := c.ItemHandler
	catHandler := c.CatHandler

	//查询cat
	catList := make([]models.Category, 0)
	catHandler.List(&catList)

	maxSize := 4

	for _, cat := range catList {
		//查询item(cat/4条)
		itemVO := vo.SubjectVO{}
		itemList := make([]models.Item, 0)
		itemHandler.ListByCat(cat.ID, maxSize, &itemList)
		itemVO.Items = itemList
		itemVO.CatID = cat.ID
		itemVO.CatTitle = cat.Title
		itemVO.CatSubTitle = cat.SubTitle
		*vos = append(*vos, itemVO)
	}

	return c
}
