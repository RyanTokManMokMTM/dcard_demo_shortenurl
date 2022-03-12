package main

import (
	"fmt"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/model"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/internal/router"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	if err := setUpSetting(); err != nil {
		log.Fatalln(err)
	}

	if err := setUpDatabase(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	//Create a http server
	server := http.Server{
		//http server setting
		Addr:           fmt.Sprintf("%s:%s", global.ServerSetting.Host, global.ServerSetting.Port), //server host
		Handler:        setUpServer(),                                                              //engine included http interface ServeHTTP
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func setUpServer() *gin.Engine {
	engine := gin.New()
	router.NewRouter(engine)
	return engine
}

func setUpSetting() error {
	//load and read the configure file
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}

	//value of the kay and store at the interface
	if err = set.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}

	if err = set.ReadSection("Database", &global.DBSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func setUpDatabase() error {
	var err error
	//set global database instance
	global.DB, err = model.NewEngine(global.DBSetting)

	if err != nil {
		return err
	}
	return nil
}
