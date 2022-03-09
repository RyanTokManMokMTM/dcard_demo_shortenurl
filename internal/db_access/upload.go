package db_access

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/model"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/service"
)

//UploadURL create the record by param
func (dao *DAO) UploadURL(param *service.UrlUploadReq) (*model.UploadModel, error) {
	upload := model.UploadModel{
		OriginalURL: param.Url,
		ExpiredAt:   param.ExpiredTime,
	}
	return upload.CreateShortenURL(dao.db)
}
