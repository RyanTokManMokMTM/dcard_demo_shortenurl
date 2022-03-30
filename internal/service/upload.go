package service

import (
	"fmt"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/util"
	"time"
)

type UrlUploadReq struct {
	//URL
	URL string `form:"url" json:"url" binding:"required,url"`
	//Expired time
	ExpiredTime time.Time `form:"ExpiredTime" json:"ExpiredTime" binding:"required,gt"`
}

type ShortenURLInfo struct {
	Id       string
	ShortUrl string
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
		Id:       shortenStr,
		ShortUrl: fmt.Sprintf("http://%s:%s/%s", global.ServerSetting.Host, global.ServerSetting.Port, shortenStr),
	}, err
}
