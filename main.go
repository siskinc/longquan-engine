package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/longquan-engine/global"
	"github.com/siskinc/longquan-engine/routes"
)

// @title Swagger Example API
// @version 1.0
// @description 龙泉决策引擎
// @termsOfService http://swagger.io/terms/

// @contact.name Daryl
// @contact.url https://www.daryl.top/
// @contact.email susecjh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /engine
func main() {
	router := gin.Default()
	routes.Init(router)
	router.Run(global.Config.ServerAddress())
}
