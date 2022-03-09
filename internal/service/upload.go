package service

import "time"

type UrlUploadReq struct {
	//URL
	Url string `form:"url" binding:"required,url"`
	//Expired time
	ExpiredTime time.Duration `form:"expired_time" binding:"required"`
}

type ShortenURLInfo struct {
	ShortenURL string
	LongestURL string
	ExpiredAty time.Duration
}

func (serve *Service) CreateShortenUrl(param *UrlUploadReq) (*ShortenURLInfo, error) {
	//calling dao
	//process the upload service
	_, err := serve.dao.UploadURL(param)
	if err != nil {
		return nil, err
	}

	//generate the shortenURL by base62 base on time or id
	return nil, nil
}
