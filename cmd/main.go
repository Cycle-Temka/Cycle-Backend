package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	runErr := server.Run("localhost:8000")

	if runErr != nil {
		panic(runErr)
	}
}
