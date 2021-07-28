package config

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) {
	configGrouper := router.Group("/config")
	configGrouper.GET("/property-type-enum", PropertyTypeEnumData)
	configGrouper.GET("/property-system-class-enum", PropertySystemClassEnumData)
}
