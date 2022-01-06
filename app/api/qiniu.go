package api

import (
	"github.com/gin-gonic/gin"
	"gomo/common/apis"
	"gomo/qiniu"
)

type Qiniu struct {
	apis.Api
}

func (q Qiniu) GetUpToken(ctx *gin.Context) {
	token := qiniu.GetToken()
	q.OK(token, "ok")
}

