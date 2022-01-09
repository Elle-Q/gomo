package api

import (
	"github.com/gin-gonic/gin"
	"gomo/app/service/dto"
	"gomo/common/apis"
	"gomo/db/handlers"
	"gomo/db/models"
	"strconv"
	"strings"
)

type Config struct {
	apis.Api
}

//查询系统默认头像配置
func (e Config) FindDefaultAvatarByName(ctx *gin.Context) {
	req := dto.ConfigAvatarApiReq{}
	configService := handlers.ConfigHandler{}
	fileService := handlers.FileHandler{}

	err := e.MakeContext(ctx).
		MakeDB().
		Bind(&req, nil).
		MakeService(&configService.Handler).
		MakeService(&fileService.Handler).
		Errors

	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	//找到配置项
	var config models.Config
	err = configService.FindByName(req.Name, &config).Error
	if err != nil {
		e.Error(500, err, "fail")
		return
	}

	vals := strings.Split(config.Val, ",")
	ids := make([]int, len(vals))
	for i := range ids {
		ids[i], _ = strconv.Atoi(vals[i])
	}

	//读取配置项值, 查找file表
	var files []models.File
	fileService.QueryByIds(config.Val, &files)

	var avatars []string
	for _, f := range files {
		avatars = append(avatars, f.QnLink)
	}
	e.OK(avatars, "ok")


}
