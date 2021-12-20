package dto

import (
	"errors"
	"gomo/db/models"
	"mime/multipart"
	"time"
)

type CatApiReq struct {
	Id int `uri:"id"`
}

func (s *CatApiReq) GetId() int {
	return s.Id
}

type CatUpdateReq struct {
	ID       int    `form:"ID" comment:"id"`        // 编码
	Title    string `json:"Title" comment:"标题"`     //菜单name
	SubTitle string `form:"SubTitle" comment:"副标题"` //菜单name
	Preview  *multipart.FileHeader  `form:"Preview" comment:"主图"`   //菜单name
	Desc     string `form:"Desc" comment:"描述"`      //菜单name
	Status   string `form:"Status" comment:"状态"`    //菜单name
}

func (s *CatUpdateReq) Generate(model *models.Category) error{

	err := s.check()
	if err != nil {
		return err
	}

	if s.ID != 0 {
		model.ID = s.ID
	}
	model.Title = s.Title
	model.SubTitle = s.SubTitle
	if s.Preview != nil {
		model.Preview = s.Preview.Filename
	}
	model.Desc = s.Desc
	model.Status = s.Status
	model.UpdateTime = time.Now()
	model.CreateTime = time.Now()

	return nil
}

func (s *CatUpdateReq) check() error{
	if len(s.Title) < 1 || len(s.SubTitle) < 1 || len(s.Desc) < 1 || len(s.Status) < 1 {
		return errors.New("参数不能为空")
	}
	return nil
}


type CatDeleteApiReq struct {
	ID int `json:"id"`
}

func (s *CatDeleteApiReq) GetId() int {
	return s.ID
}