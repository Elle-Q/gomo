package api

import (
	"github.com/gin-gonic/gin"
	"gomo/tool"
	"log"
)

func getUserBuyId(id int, ctx *gin.Context)  {

	db, err := tool.GetDB(ctx)
	if err != nil {
		log.Fatal("db not found")
	}

	db.Query("select * from user where id = 1")
}