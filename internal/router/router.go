package router

import (
	v1 "github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/api/v1"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

//NewRouter all router define here
func NewRouter(engine *gin.Engine) {
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type"},
		AllowCredentials: true,
	}))
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Msg": "Hello axis",
		})
	})
	engine.Use(middleware.RateLimiter())
	engine.Use(middleware.ValidateTranslator())

	apiV1 := engine.Group("/api/v1")
	apiV1.GET("/limiter", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "succeed",
		})
	})
	apiV1.POST("/urls", v1.UploadController)
	engine.GET("/:url_id", v1.GetUrlAndRedirect)
	if gin.Mode() == gin.DebugMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	}
}
