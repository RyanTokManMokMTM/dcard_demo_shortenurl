package service

type UrlUploadReq struct {
	//URL
	Url string `form:"url"`
	//Expired time
	ExpiredTime string `form:"expired_time"`
}

func (serve *Service) GetShortenURL(param *UrlUploadReq) {
	//calling dao
	//process the upload service
}
