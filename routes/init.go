package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginxMiddlewares "github.com/goools/tools/ginx/middlewares"
	"github.com/siskinc/longquan-engine/global"
	namesapce "github.com/siskinc/longquan-engine/routes/namespace"
	"github.com/siskinc/longquan-engine/routes/property"
	_ "github.com/siskinc/longquan-engine/swagger_docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(router *gin.Engine) {
	router.Use(cors.Default())
	router.Use(gin.LoggerWithFormatter(ginxMiddlewares.LoggerFormatter))

	apiRouter := router.Group("/engine/api")
	namesapce.InitRouter(apiRouter)
	property.InitRouter(apiRouter)

	// swagger
	{
		url := ginSwagger.URL("http://" + global.Config.SwaggerHost + "/swagger/doc.json") // The url pointing to API definition
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
