package main

import (
	"fmt"
	"go-web-base/config"
	"go-web-base/global"
	"go-web-base/initialiaze"
	"net/http"
	"time"
)

func main() {
	config.InitZap()
	config.InitViper()
	config.InitMysql()
	initialiaze.DbTables()
	defer global.DB.Close()
	addr := fmt.Sprintf(":%d", global.CONFIG.System.Port)
	server := http.Server{
		Addr: addr,
		Handler: initialiaze.Routers(),
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
