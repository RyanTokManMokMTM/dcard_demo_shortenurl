package service

import (
	"errors"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"time"
)

func (serve *Service) ShortenInfo(urlId string) (string, error) {
	//get shorten info
	info, err := serve.dao.GetShortenURLInfo(urlId)
	if err != nil {
		return "", err
	}
	//compare with the time
	//expired time can't less than 1 or n minute ,
	today := time.Now()
	if today.Sub(info.ExpiredAt) > global.AppSetting.NotAllowedAccessTime { //suppose to day is 2-13 and expired time is 2-10 -> 3*24*60*60
		return "", errors.New("URL GetCode expired")
	}

	return info.OriginalURL, nil
}
