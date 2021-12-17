package dto

import (
	"gomo/db/models"
	"mime/multipart"
	"time"
)

type UserApiReq struct {
	Id int `uri:"id"`
}

func (s *UserApiReq) GetId() int {
	return s.Id
}

type UserUpdateApiReq struct {
	Id int `form:"id"`
	Name string `form:"Name"`
	Address string `form:"Address"`
	Gender string `form:"Gender"`
	Avatar  *multipart.FileHeader  `form:"Avatar" comment:"头像"`
	BgImag  *multipart.FileHeader  `form:"BgImag" comment:"背景"`
}

func (u *UserUpdateApiReq) Generate() models.User{
	user := models.User{}
	user.ID = u.Id
	user.Name = u.Name
	user.Address = u.Address
	user.Gender = u.Gender
	user.Avatar = u.Avatar.Filename
	user.BgImag = u.BgImag.Filename
	user.UpdateTime=time.Now()
	return user
}

type UserLoginApiReq struct {
	UserName string `form:"UserName"`
	Password string `form:"Password"`
}
