package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-gin-admin/types"
	"time"
)

// Result 返回结果
type Result struct {
	Time time.Time   `json:"time"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Time = time.Now()
	res.Code = int(types.ApiCode.SUCCESS)
	res.Msg = types.ApiCode.GetMessage(types.ApiCode.SUCCESS)
	res.Data = data

	c.JSON(http.StatusOK, res)
}

func Error(c *gin.Context, code int, msg string) {
	res := Result{}
	res.Time = time.Now()
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}

	c.JSON(http.StatusOK, res)
}
