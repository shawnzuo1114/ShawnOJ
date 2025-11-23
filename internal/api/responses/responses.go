package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Responses struct {
	Code ResCode `json:"code"`
	Msg  string  `json:"msg"`
	Data any     `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &Responses{
		Code: code,
		Msg:  code.Msg(),
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithData(c *gin.Context, code ResCode, data any) {
	rd := &Responses{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data any) {
	rd := &Responses{
		Code: CodeSuccess,
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
