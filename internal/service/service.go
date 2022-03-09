package service

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/db_access"
	"github.com/gin-gonic/gin"
)

type Service struct {
	dao *db_access.DAO
	ctx *gin.Context
}

func NewService(ctx *gin.Context) *Service {
	return &Service{ctx: ctx, dao: db_access.NewDAO(global.DB)}
}
