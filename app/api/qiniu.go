package api

import (
	"github.com/gin-gonic/gin"
	"gomo/common/apis"
	"gomo/config"
	"gomo/qiniu/img"
)

type Qiniu struct {
	apis.Api
}

func (q Qiniu) GetUpToken(ctx *gin.Context) {
	q.Context = ctx
	token := img.GetPubImgToken()
	resp := map[string]string{
		"UpToken": token,
		"Domain": config.QiniuConfig.PubDomain,
	}
	q.OK(resp, "ok")
}

