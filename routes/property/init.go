package property

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.RouterGroup) {
	propertySetGrouper := router.Group("/property-set")
	propertyGrouper := router.Group("/property")
	// 属性集
	propertySetGrouper.GET("/", QueryPropertySet)
	propertySetGrouper.POST("/", CreatePropertySet)
	propertySetGrouper.PATCH("/:id", UpdatePropertySet)
	propertySetGrouper.DELETE("/:id", DeletePropertySet)

	// 属性
	propertyGrouper.GET("/", QueryProperty)
	propertyGrouper.POST("/", CreateProperty)
	propertyGrouper.PATCH("/:id", UpdateProperty)
	propertyGrouper.DELETE("/:id", DeleteProperty)
}
