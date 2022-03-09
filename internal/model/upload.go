package model

import (
	"gorm.io/gorm"
	"time"
)

type UploadModel struct {
	*Model
	OriginalURL string
	ShortenURL  string
	ExpiredAt   time.Duration
}

func (up UploadModel) GetTableName() string {
	return "upload_model"
}

func (up UploadModel) CreateShortenURL(db *gorm.DB) (*UploadModel, error) {
	//whether longestURL is existed
	//time := time.Now()
	//check long url is existed and not expired
	//if the record is existed and expired,just updated the expired date and return the
	err := db.Model(up).Where("original_url = ?", up.OriginalURL).Error
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err = db.Create(up).Error; err != nil {
		return nil, err
	}
	return &up, err
}

func (up UploadModel) UpdateShortenURL(id int64) {}

//func (up*UploadModel) GetURLByUrlID(db *gorm.DB, urlID string) string {
//	return global.DBSetting.TablePrefix + "shortenURL"
//}
