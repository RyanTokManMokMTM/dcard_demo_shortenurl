package service

import (
	"errors"
	"time"
)

func (serve *Service) ShortenInfo(urlId string) (string, error) {
	//get shorten info
	info, err := serve.dao.GetShortenURLInfo(urlId)
	if err != nil {
		return "", err
	}

	//compare with the time
	//expired time can't less than 1 minute ,
	today := time.Now()
	if today.Sub(info.ExpiredAt).Seconds() < 60 {
		return "", errors.New("url id is closer to expired time,please update or create a new one")
	}

	return info.OriginalURL, nil
}
