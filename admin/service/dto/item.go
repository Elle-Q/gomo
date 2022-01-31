package dto

import (
	"errors"
)

type ItemUpdateReq struct {
	ID      int  `json:"ID" comment:"id"`       // id
	CatId   int64  `json:"CatId" comment:"CatId"` // cat id
	Name    string `json:"Name" comment:"名称"`     //名称
	Tags    []string `json:"Tags" comment:"标签"`     //名称
	BLink   string `json:"BLink" comment:"B站链接"`  //B站链接
	Preview string `json:"Preview" comment:"预览图"` //预览图
	Desc    string `json:"Desc" comment:"描述"`     //描述
	Author  string `json:"Author" comment:"作者"`   //状态
	Price   int  `json:"Price" comment:"价格"`    //状态
	Status  string `json:"Status" comment:"状态"`   //状态
}

func (s ItemUpdateReq) check() error {
	if len(s.Name) < 1 || len(s.Preview) < 1 || len(s.Desc) < 1 || len(s.Author) < 1 || len(s.Status) < 1 {
		return errors.New("参数不能为空")
	}
	return nil
}
