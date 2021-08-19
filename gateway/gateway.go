package main

import (
	"github.com/gin-gonic/gin"
	h "github.com/petrolax/grpc-test/gateway/handler"
)

func main() {
	router := gin.Default()

	router.POST("/hello", h.Hello)
	router.POST("/bye", h.Bye)

	router.Run(":8080")
}
