package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// DeleteProperty godoc
// @TAGS 场景管理
// @Summary 删除属性
// @Description 删除属性
// @Accept json
// @Produce json
// @Param id path string true "属性id" minlength(1)
// @Success 200 {object} httpx.JSONResult
// @Router /property/{id} [delete]
func DeleteProperty(c *gin.Context) {
	propertyId := c.Param("id")
	service := propertyService.NewPropertyService()
	err := service.DeleteOne(propertyId)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
}
