package dto

import (
	"gomo/db/models"
	"time"
)

type UserApiReq struct {
	Id int `uri:"id"`
}

func (s *UserApiReq) GetId() int {
	return s.Id
}

type UserUpdateApiReq struct {
	Id int `json:"id"`
	Name string `json:"Name"`
	Address string `json:"Address"`
	Gender string `json:"Gender"`
	Moto string `json:"Moto"`
	Status string `json:"Status"`
	//Avatar  *multipart.FileHeader  `form:"Avatar" comment:"头像"`
	//BgImag  *multipart.FileHeader  `form:"BgImag" comment:"背景"`
}

func (u *UserUpdateApiReq) Generate() models.User{
	user := models.User{}
	user.ID = u.Id
	user.Name = u.Name
	user.Address = u.Address
	user.Gender = u.Gender
	user.Moto = u.Moto
	user.Status = u.Status
	user.UpdateTime=time.Now()
	return user
}

type UserLoginApiReq struct {
	UserName string `form:"UserName"`
	Password string `form:"Password"`
}

type UserTokenRefreshApiReq struct {
	RefreshToken string `form:"RefreshToken"`
}

type UserUpdateAvatarApiReq struct {
	ID int `json:"UserId"`
	Avatar  string  `json:"Avatar" comment:"头像"`
}

type UserUpdateBGApiReq struct {
	ID int `json:"UserId"`
	BgImag  string  `json:"BgImag" comment:"背景"`
}