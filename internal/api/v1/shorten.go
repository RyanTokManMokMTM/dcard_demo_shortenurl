package v1

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/service"
	//_ "github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/errCode"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/app"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/errCode"
	"github.com/gin-gonic/gin"
)

/*
GetUrlAndRedirect -TODO
-Checking stage
	need to get the url code if from query parameter
	query database
		-whether url code is exist
		-whether url code is expired
- Response
	redirect
*/

// @Summary Redirect to original url by url id
// @Tags RedirectURL
// @Version 1.0
// @Produce text/html
// @Param url_id path string required "short url id"
// @Success 301 string string "redirect to orginal url"
// @Failure 404 {obejct} errCode.Error "url id expired/not existed"
// @Router /{url_id} [GET]
func GetUrlAndRedirect(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	id := ctx.Param("url_id")
	//if err != nil {
	//	res.ErrorResponse(errCode.InvalidParams.WithDetail(err.Error()))
	//	return
	//}

	serve := service.NewService(ctx)
	originalURL, err := serve.ShortenInfo(id)
	if err != nil {
		res.ErrorResponse(errCode.NotFound.WithDetail(err.Error()))
		return
	}

	res.SuccessAndRedirectPermanently(originalURL)
	return
}
