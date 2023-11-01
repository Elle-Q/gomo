package dto

import (
	"errors"
)

type ItemUpdateReq struct {
	ID     int     `json:"ID" comment:"id"`       // id
	CatId  int64   `json:"CatId" comment:"CatId"` // cat id
	Name   string  `json:"Name" comment:"名称"`     //名称
	Type   string  `json:"Type" comment:"类型"`     //资源类型
	Tags   string  `json:"Tags" comment:"标签"`     //标签
	BLink  string  `json:"BLink" comment:"B站链接"`  //B站链接
	Desc   string  `json:"Desc" comment:"描述"`     //描述
	Author string  `json:"Author" comment:"作者"`   //作者
	Price  float64 `json:"Price" comment:"价格"`    //价格
	Status string  `json:"Status" comment:"状态"`   //状态
}

type ItemRescUploadReq struct {
	ItemID int    `form:"ItemID" comment:"item_id"` // item_id
	Type   string `form:"Type" comment:"资源类型"`      // Type (main, preview, refs)
}

type ItemChapterUploadReq struct {
	ItemID    int `form:"ItemID" comment:"item_id"` // item_id
	ChapterID int `form:"ChapterID" comment:"章节id"` // Type (main, preview, refs)
}

func (s ItemUpdateReq) check() error {
	if len(s.Name) < 1 || len(s.Desc) < 1 || len(s.Author) < 1 || len(s.Status) < 1 {
		return errors.New("参数不能为空")
	}
	return nil
}

type ItemIDReq struct {
	ID int `uri:"ID" comment:"ID"` // item_id
}

type ItemFileDelReq struct {
	FileId int    `json:"FileId" comment:"文件id"` // FileId
	Bucket string `json:"Bucket" comment:"七牛空间"` // Bucket
	Key    string `json:"Key" comment:"七牛key"`   // Key
}

type ChapterFileDelReq struct {
	ChapterId int    `json:"ChapterId" comment:"章节id"` // FileId
	FileId    int    `json:"FileId" comment:"文件id"`    // FileId
	Bucket    string `json:"Bucket" comment:"七牛空间"`    // Bucket
	Key       string `json:"Key" comment:"七牛key"`      // Key
	Type      string `json:"Type" comment:"类型"`        // Type
}
