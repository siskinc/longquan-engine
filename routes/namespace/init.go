package namesapce

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.RouterGroup) {
	grouper := router.Group("/namespace")
	// TODO 获取命名空间信息
	// grouper.GET("/", CreateNamespace)
	// TODO 创建命名空间
	grouper.POST("/", CreateNamespace)
	// // TODO 修改命名空间
	// grouper.PATCH("/:id")
	// TODO 删除命名空间
	grouper.DELETE("/:id", DeleteNamespace)
}
