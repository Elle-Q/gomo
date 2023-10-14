package api

import (
	"github.com/gin-gonic/gin"
	"leetroll/app/service/dto"
	"leetroll/common/apis"
	"leetroll/db/handlers"
	"leetroll/db/models"
	"leetroll/qiniu"
	"strconv"
	"strings"
)

type Config struct {
	apis.Api
}

// 查询系统默认头像配置
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
	err2 := fileService.QueryByIds(config.Val, &files).Error

	if err2 != nil {
		e.Error(500, err2, err2.Error())
		return
	}

	var avatars []string
	for _, f := range files {
		avatars = append(avatars, qiniu.GetPubUrl(f.Key))
	}
	e.OK(avatars, "ok")

}
