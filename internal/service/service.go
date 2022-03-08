package service

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	db  *gorm.DB
	ctx *gin.Context
}

func NewService(ctx *gin.Context) *Service {
	return &Service{ctx: ctx, db: global.DB}
}
