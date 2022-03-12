package v1

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/service"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/app"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/errCode"
	"github.com/gin-gonic/gin"
)

func UploadController(ctx *gin.Context) {
	//TODO - Get Longest URL by url id if id is existed
	res := app.NewResponse(ctx)
	param := service.UrlUploadReq{}
	valid, errs := app.BindingAndValidating(ctx, &param)
	if !valid {
		res.ErrorResponse(errCode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	serve := service.NewService(ctx)
	data, err := serve.CreateShortenUrl(&param)
	if err != nil {
		res.ErrorResponse(errCode.ErrorCreateShortenURL.WithDetail(err.Error()))
		return
	}

	res.SuccessResponse(data)
	return
}
