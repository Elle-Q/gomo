package service

import (
	"gomo/app/service/dto"
	"gomo/common/services"
	"gomo/db/handlers"
	"gomo/db/models"
)

type UserService struct {
	UserHandler  *handlers.UserHandler
	services.Service
}

func (s *UserService) Update(m *models.User) *UserService {
	if err := s.UserHandler.Update(m).Error; err != nil {
		_=s.AddError(err)
	}
	return s
}

func (s *UserService) Login(u *dto.UserLoginApiReq) *UserService{
	return nil
}

func (s *UserService) FindById(u *dto.UserApiReq, m *models.User) *UserService {
	if err := s.UserHandler.FindById(u.GetId(), m).Error; err != nil {
		_=s.AddError(err)
	}
	return s
}

