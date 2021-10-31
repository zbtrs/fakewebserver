package main

import (
	"fakewebserver/fakewebserver/src/route"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := route.RegisterWebRoute()
	router.Run(":8080")
}
