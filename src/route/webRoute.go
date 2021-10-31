package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegisterWebRoute() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.POST("/github.com/session", controllers.SessionHandler)
	router.GET("/github.com/login", controllers.LoginHandler)
	router.GET("/github.com", controllers.GithubHandler)

	return router
}
