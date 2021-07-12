package namesapce

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.RouterGroup) {
	grouper := router.Group("/namespace")
	// 获取命名空间信息
	grouper.GET("/", QueryNamespace)
	// 创建命名空间
	grouper.POST("/", CreateNamespace)
	// 修改命名空间
	grouper.PATCH("/:id", UpdateNamespace)
	// 删除命名空间
	grouper.DELETE("/:id", DeleteNamespace)
}
