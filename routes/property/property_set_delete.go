package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// DeletePropertySet godoc
// @TAGS 场景管理
// @Summary 删除属性集
// @Description 删除属性集
// @Accept json
// @Produce json
// @Param id path string true "属性集id" minlength(1)
// @Success 200 {object} httpx.JSONResult
// @Router /property-set/{id} [delete]
func DeletePropertySet(c *gin.Context) {
	propertySetId := c.Param("id")
	service := propertyService.NewPropertySetService()
	err := service.Delete(propertySetId)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
}
