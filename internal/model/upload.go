package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type UploadModel struct {
	*Model
	OriginalURL string
	ShortenURL  string
	ExpiredAt   time.Time
}

func (up UploadModel) GetTableName() string {
	return "upload_model"
}

func (up UploadModel) CreateShortenURL(db *gorm.DB) (*UploadModel, error) {

	//whether longestURL is existed
	//time := time.Now()
	//check long url is existed and not expired
	//if the record is existed and expired,just updated the expired date and return the
	var exist int64
	db.Model(&up).Select("count(*)").Where("original_url = ?", up.OriginalURL).Find(&exist)
	log.Println(exist)
	if exist > 0 {
		return nil, errors.New("original_url is already exist")
	}

	if err := db.Create(&up).Error; err != nil {
		return nil, err
	}
	return &up, nil
}

func (up UploadModel) UpdateShortenURL(db *gorm.DB, v interface{}) error {
	err := db.Model(&up).Where("id = ? AND original_url =?", up.ID, up.OriginalURL).Updates(v).Error
	return err
}

func (up UploadModel) GetShortenURLInfo(db *gorm.DB) (*UploadModel, error) {
	if err := db.Model(&up).Where("shorten_url = ?", up.ShortenURL).First(&up).Error; err != nil {
		return nil, err
	}

	return &up, nil
}
