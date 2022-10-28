package app

import "github.com/gin-gonic/gin"

func InitGinEngine() *gin.Engine {
	app := gin.New()
	return app
}
