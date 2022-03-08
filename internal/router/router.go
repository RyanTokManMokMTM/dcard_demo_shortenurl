package router

import (
	v1 "github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

//NewRouter all router define here
func NewRouter(engine *gin.Engine) {
	apiv1 := engine.Group("/api/v1")

	apiv1.POST("/urls", v1.UploadURLs)
	//apiv1.GET("/:url_id")
	engine.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "testing",
		})
		return
	})
}
