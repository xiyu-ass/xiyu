package main

import (
	"myproject3/config"
	"myproject3/routes"
	"myproject3/utils/logger"
)

func main() {

	db := config.InitDB()
	defer db.Close()

	initGin() //初始化gin框架并启动

}

func initGin() {
	r := routes.Init()
	err := r.Run()
	if err != nil {
		logger.PanicError(err, "Server startup failed!", true)
	}
}
