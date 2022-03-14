package v1

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/service"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/app"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/errCode"
	"github.com/gin-gonic/gin"
)

// @Summary Upload longest url with expired time
// @Tags UploadURL
// @Version 1.0
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param URL formData string true "original url"
// @Param ExpiredTime formData string true "UTC time"
// @Success 200 {object} service.ShortenURLInfo "upload succeed"
// @Failure 500 {object} errCode.Error "upload failed"
// @Failure 400 {object} errCode.Error "request parameter invaild"
// @Router /api/v1/urls [POST]
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
