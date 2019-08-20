package main

import (
	"fmt"
	"github.com/drip/hello/pkg/setting"
	"github.com/drip/hello/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

// Initialization operation
func init() {
	setting.Setup()
}

func main() {
	// Loading router
	router := routers.InitRouter()

	// Loading static resources
	router.Static("/asset", "./asset")

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
