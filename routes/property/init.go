package property

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.RouterGroup) {
	propertySetGrouper := router.Group("/property-set")
	propertyGrouper := router.Group("/property")
	// 属性集
	propertySetGrouper.GET("/")
	propertySetGrouper.POST("/")
	propertySetGrouper.PATCH("/:id")
	propertySetGrouper.DELETE("/:id")

	// 属性
	propertyGrouper.GET("/")
	propertyGrouper.POST("/")
	propertyGrouper.PATCH("/:id")
	propertyGrouper.DELETE("/:id")
}
