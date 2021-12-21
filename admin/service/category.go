package service

import (
	"gomo/app/service/dto"
	"gomo/common/services"
	"gomo/db/handlers"
	"gomo/db/models"
)

type CategoryService struct {
	CatHandler  *handlers.CatHandler
	services.Service
}

func (s *CategoryService) List(models *[]models.Category) *CategoryService{
	if err := s.CatHandler.List(models).Error; err != nil {
		_=s.AddError(err)
	}
	return s
}

func (s *CategoryService) Get(req *dto.CatApiReq, cat *models.Category) *CategoryService{
	if err := s.CatHandler.Get(req.GetId(), cat).Error; err != nil {
		_=s.AddError(err)
	}
	return s
}

