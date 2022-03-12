package router

import (
	v1 "github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/api/v1"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/middleware"
	"github.com/gin-gonic/gin"
)

//NewRouter all router define here
func NewRouter(engine *gin.Engine) {

	//engine.Use(middleware.ValidateCustomFields())
	engine.Use(middleware.ValidateTranslator())

	apiV1 := engine.Group("/api/v1")

	apiV1.POST("/urls", v1.UploadController)
	engine.GET("/:url_id", v1.GetUrlAndRedirect)

}
