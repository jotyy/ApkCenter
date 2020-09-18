package main

import (
	"file-upload/conf"
	"file-upload/router"
	"log"
)

// @title 和行APK管理中心
// @version 1.0
// @description 方便快捷的APK管理界面.
func main() {
	conf.Init()

	routerInit := router.NewRouter()

	err := routerInit.Run("0.0.0.0:8686")

	//server := &http.Server{
	//	Addr:              ":8686",
	//	Handler:           routerInit,
	//	ReadTimeout:       60,
	//	WriteTimeout:      60,
	//	MaxHeaderBytes:    1 << 20,
	//}
	//
	//log.Printf("[info] start http server listening %s", ":8686")
	//
	//server.ListenAndServe()

	//endless.DefaultReadTimeOut = 60
	//endless.DefaultWriteTimeOut = 60
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//server := endless.NewServer(":8686", routerInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
