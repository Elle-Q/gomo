package apis

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"leetroll/common/response"
	"leetroll/db"
	"leetroll/tool"
)

type Api struct {
	Context *gin.Context
	DB      *sql.DB
	Errors  error
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else {
		e.Errors = fmt.Errorf("%v;%w", e.Error, err)
	}
}

func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	return e
}

func (e Api) GetDB() (*sql.DB, error) {
	db, err := tool.GetDB(e.Context)
	if err != nil {
		e.Error(500, err, "数据库连接获取失败")
		return nil, err
	}
	return db, nil
}

func (e *Api) MakeDB() *Api {
	db, err := tool.GetDB(e.Context)
	if err != nil {
		e.AddError(err)
	}
	e.DB = db
	return e
}

func (e *Api) MakeService(handler *db.Handler) *Api {
	handler.DB = e.DB
	return e
}

func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = append(bindings, binding.JSON, nil)
	}
	for i := range bindings {
		switch bindings[i] {
		case binding.JSON:
			err = e.Context.ShouldBindWith(d, binding.JSON)
		case binding.XML:
			err = e.Context.ShouldBindWith(d, binding.XML)
		case binding.Form:
			err = e.Context.ShouldBindWith(d, binding.Form)
		case binding.Query:
			err = e.Context.ShouldBindWith(d, binding.Query)
		case binding.FormPost:
			err = e.Context.ShouldBindWith(d, binding.FormPost)
		case binding.FormMultipart:
			err = e.Context.ShouldBindWith(d, binding.FormMultipart)
		case binding.ProtoBuf:
			err = e.Context.ShouldBindWith(d, binding.ProtoBuf)
		case binding.MsgPack:
			err = e.Context.ShouldBindWith(d, binding.MsgPack)
		case binding.YAML:
			err = e.Context.ShouldBindWith(d, binding.YAML)
		case binding.Header:
			err = e.Context.ShouldBindWith(d, binding.Header)
		default:
			err = e.Context.ShouldBindUri(d)
		}
		if err != nil {
			e.AddError(err)
		}
	}
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	response.Custum(e.Context, data)
}
