package property

import "github.com/gin-gonic/gin"

// PropertySerMapQuery godoc
// @TAGS 场景管理
// @Summary 查询属性地图
// @Description 查询属性地图
// @Accept json
// @Produce json
// @Param id path string true "属性集id" minlength(1)
// @Success 200 {object} httpx.JSONResultPaged.{data=[]propertyModel.PropertySet}
// @Router /namespace-set [get]
func PropertySerMapQuery(c *gin.Context) {

}
