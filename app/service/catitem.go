package service

import (
	"leetroll/app/service/vo"
	"leetroll/db/handlers"
	"leetroll/db/models"
	"leetroll/qiniu"
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
		subjectVo := vo.SubjectVO{}
		itemList := make([]models.Item, 0)
		itemHandler.ListPopularByCat(cat.ID, maxSize, &itemList)
		for i, _ := range itemList {
			itemList[i].Main = qiniu.GetPrivateUrl(itemList[i].Main)
		}
		subjectVo.Items = itemList
		subjectVo.CatID = cat.ID
		subjectVo.CatTitle = cat.Title
		subjectVo.CatSubTitle = cat.SubTitle
		*vos = append(*vos, subjectVo)
	}

	return c
}
