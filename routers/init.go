package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginxMiddlewares "github.com/goools/tools/ginx/middlewares"
)

func Init(router *gin.Engine) {
	router.Use(cors.Default())
	router.Use(gin.LoggerWithFormatter(ginxMiddlewares.LoggerFormatter))
}
