package app

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/errCode"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Response format

type Response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx: ctx}
}

func (res *Response) ErrorResponse(err *errCode.Error) {
	//response message
	data := gin.H{
		"code": err.GetCode(),
		"msg":  err.GetMsg(),
	}

	detail := err.GetDetail()
	if len(detail) > 0 {
		data["detail"] = detail
	}

	res.ctx.JSON(err.StatusCode(), data)
}

func (res *Response) SuccessResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	res.ctx.JSON(http.StatusOK, data)
}

func (res *Response) SuccessAndRedirectPermanently(url string) {
	res.ctx.Redirect(http.StatusMovedPermanently, url) //move to url permanently
}
