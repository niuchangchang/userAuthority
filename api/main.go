package main

import (
	"github.com/wangcong0918/sunrise"
	"userAuthority/api/middleware"
	"userAuthority/api/routers"
	"os"
	"time"
)

var PORT string
var cstZone = time.FixedZone("CST", 8*3600)

func init() {
	time.LoadLocation("Asia/Shanghai")
	time.Local = cstZone
	PORT = os.Getenv("PORT")
	if PORT == "" {
		panic("http point not nil")
	}
}

func main() {
	router := sunrise.Default()
	router.Use(middleware.Cors)
	routers.User(router)
	//routers.Device(router)
	//routers.Area(router)
	 router.Run(PORT)
}
