package tool

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
)

func  GetDB(c *gin.Context) (*sql.DB, error) {

	idb, exist := c.Get("db")
	if !exist {
		return nil, errors.New("db connect not exist")
	}
	switch idb.(type) {
	case *sql.DB:
		return idb.(*sql.DB), nil
	default:
		return nil, errors.New("db connect not exist")
	}


}
