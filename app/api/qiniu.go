package api

import (
	"github.com/gin-gonic/gin"
	"leetroll/common/apis"
	"leetroll/config"
	"leetroll/qiniu/regular"
)

type Qiniu struct {
	apis.Api
}

func (q Qiniu) GetPubUpToken(ctx *gin.Context) {
	q.Context = ctx
	token := regular.GetToken(config.QiniuConfig.PubBucket)
	resp := map[string]string{
		"UpToken": token,
		"Domain":  config.QiniuConfig.PubDomain,
	}
	q.OK(resp, "ok")
}
