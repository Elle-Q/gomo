package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Default = &response{}

//响应失败处理
func Error(c *gin.Context, code int, err error, msg string)  {
	res := Default.Clone()
	if err != nil {
		res.SetMsg(err.Error())
	}

	if msg != "" {
		res.SetMsg(msg)
	}

	res.SetCode(code)
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(http.StatusOK, res)
}


// 相应成功处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

// Custum 兼容函数
func Custum(c *gin.Context, data gin.H) {
	c.Set("result", data)
	c.AbortWithStatusJSON(http.StatusOK, data)
}