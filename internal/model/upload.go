package model

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"gorm.io/gorm"
	"time"
)

type UploadModel struct {
	*Model
	OriginalURL string
	ShortenURL  string
	ExpiredAt   time.Duration
}

func (up *UploadModel) GetTableName() string {
	return global.DBSetting.TablePrefix + "shortenURL"
}

func (up *UploadModel) CreateShortenURL(db *gorm.DB) string {
	return global.DBSetting.TablePrefix + "shortenURL"
}

func (up *UploadModel) GetURLByUrlID(db *gorm.DB, urlID string) string {
	return global.DBSetting.TablePrefix + "shortenURL"
}
