package setting

import "time"

type (
	Server struct {
		Host         string
		Port         string
		ReadTimeOut  time.Duration
		WriteTimeOut time.Duration
		Mode         string
	}

	DB struct {
		User         string
		Password     string
		Host         string
		DBName       string
		TablePrefix  string
		Charset      string
		ParseTime    bool
		MaxIdleConns int
		MaxOpenConns int
	}

	App struct {
		NotAllowedAccessTime time.Duration
	}
)
