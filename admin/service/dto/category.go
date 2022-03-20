package dto

import (
	"errors"
	"gomo/db/models"
	"time"
)

type CatApiReq struct {
	Id int `uri:"id"`
}

func (s *CatApiReq) GetId() int {
	return s.Id
}

type CatUpdateReq struct {
	ID       int    `json:"ID" comment:"id"`        // id
	Title    string `json:"CatTitle" comment:"标题"`     //标题
	SubTitle string `json:"SubTitle" comment:"副标题"` //副标题
	Preview  string `json:"Preview" comment:"预览图"`   //预览图
	PageImg  string `json:"PageImg" comment:"主图"`   //主图
	Desc     string `json:"Desc" comment:"描述"`      //描述
	Status   string `json:"Status" comment:"状态"`    //状态
}

func (s *CatUpdateReq) Generate(model *models.Category) error{

	err := s.check()
	if err != nil {
		return err
	}

	if s.ID != 0 {
		model.ID = int64(s.ID)
	}
	model.Title = s.Title
	model.SubTitle = s.SubTitle
	model.Preview = s.Preview
	model.PageImg = s.PageImg
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

func (s *CatDeleteApiReq) GetDeleteId() int {
	return s.ID
}

type CatGetApiReq struct {
	ID int `uri:"id"`
}

func (s *CatGetApiReq) GetId() int {
	return s.ID
}