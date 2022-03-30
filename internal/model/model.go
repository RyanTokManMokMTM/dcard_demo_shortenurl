package model

import (
	"fmt"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Model Sharing model property
type Model struct {
	ID uint32 `json:"id" gorm:"primary_key"`
	//CreateOn  uint32 `json:"create_on" gorm:"NOT NULL"`
	//DeleteOn  uint32 `json:"delete_on"`
	//IsExpired int8 `json:"is_expired" gorm:"NOT NULL"`
}

//NewEngine init the engine
func NewEngine(dbSetting *setting.DB) (*gorm.DB, error) {

	set := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.User,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	)

	db, err := gorm.Open(mysql.Open(set))
	if err != nil {
		return nil, err
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = migration(db)
	if err != nil {
		return nil, err
	}

	sql.SetMaxOpenConns(dbSetting.MaxOpenConns)
	sql.SetMaxIdleConns(dbSetting.MaxIdleConns)
	return db, nil
}

func migration(db *gorm.DB) error {
	err := db.AutoMigrate(&UploadModel{})
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	if global.DB != nil {
		db, err := global.DB.DB()
		if err != nil {
			return
		}

		db.Close()
	}
}
