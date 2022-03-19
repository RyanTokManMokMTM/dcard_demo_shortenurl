package middleware

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/app"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/errCode"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"sync"
)

var (
	onceTask = sync.Once{}
)

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIP := ctx.ClientIP()
		res := app.NewResponse(ctx)
		if clientIP == "" {
			res.ErrorResponse(errCode.ClientError.WithDetail("Client agent info not found or error"))
			ctx.Abort()
		}
		//log.Println(clientIP)
		//limiter : 10 request for each user and a token will generate after 1s
		l := newLimiters(
			//1s to generate a token
			rate.Every(global.AppSetting.LimiterTokenTime),
			global.AppSetting.LimiterBucketSize, //there are total 10 buckets
			clientIP)

		if !l.Allow() {
			res.ErrorResponse(errCode.TooManyRequest)
			ctx.Abort()
		}
		ctx.Next()
	}
}

func newLimiters(r rate.Limit, b int, key string) *limiter.Limiter {
	onceTask.Do(func() {
		log.Println("run once")
		go global.Limiters.ClearNotUseLimiter(global.AppSetting.LimterClearTime)
	})
	return global.Limiters.GetLimiter(r, b, key)
}
