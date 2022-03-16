package global

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/limiter"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	Limiters *limiter.Limiters
)
