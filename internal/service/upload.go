package service

import (
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/util"
	"time"
)

type UrlUploadReq struct {
	//URL
	URL string `form:"URL" binding:"required,url"`
	//Expired time
	ExpiredTime time.Time `form:"ExpiredTime" binding:"required,gt"`
}

type ShortenURLInfo struct {
	ShortenURL string
	LongestURL string
	ExpiredAt  time.Time
}

func (serve *Service) CreateShortenUrl(param *UrlUploadReq) (*ShortenURLInfo, error) {
	//calling dao
	//process the upload service
	model, err := serve.dao.UploadURL(param.URL, param.ExpiredTime)
	if err != nil {
		return nil, err
	}

	//generate the shortenURL by base62 base on time or id
	unixTime := time.Now().Unix()
	shortenStr := util.Base62URL(unixTime + int64(model.ID))

	err = serve.dao.UpdateShortenURL(model.ID, model.OriginalURL, shortenStr)
	if err != nil {
		return nil, err
	}

	return &ShortenURLInfo{
		ShortenURL: shortenStr,
		LongestURL: model.OriginalURL,
		ExpiredAt:  model.ExpiredAt,
	}, err
}
