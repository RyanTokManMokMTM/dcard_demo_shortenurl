package db_access

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/model"
	"time"
	//"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/service"
	//"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/service"
)

//UploadURL create the record by param
func (dao *DAO) UploadURL(url string, expiredTime time.Time) (*model.UploadModel, error) {
	upload := model.UploadModel{
		OriginalURL: url,
		ExpiredAt:   expiredTime,
	}
	return upload.CreateShortenURL(dao.db)
	//return nil, nil
}

func (dao *DAO) UpdateShortenURL(id uint32, originalURL, shortenURL string) error {
	upload := model.UploadModel{Model: &model.Model{
		ID: id,
	}, OriginalURL: originalURL}

	value := map[string]interface{}{
		"shorten_url": shortenURL,
	}
	return upload.UpdateShortenURL(dao.db, value)
}

func (dao *DAO) GetShortenURLInfo(shortenURLID string) (*model.UploadModel, error) {
	upload := model.UploadModel{ShortenURL: shortenURLID}
	return upload.GetShortenURLInfo(dao.db)
}
