package main

import (
	"learning-gin/pkg/config"
	"learning-gin/pkg/database"
)

func main() {
	config.InitConfig()

	database.ConnectDB()

}
