package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// UpdateProperty godoc
// @TAGS 场景管理
// @Summary 修改属性
// @Description 修改属性
// @Accept json
// @Produce json
// @Param id path string true "属性id" minlength(1)
// @Param message body propertyService.UpdatePropertyReq true "属性更新信息"
// @Success 200 {object} httpx.JSONResult.{data=propertyModel.Property}
// @Router /property/{id} [patch]
func UpdateProperty(c *gin.Context) {
	propertyId := c.Param("id")
	req := &propertyService.UpdatePropertyReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to UpdatePropertyReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	service := propertyService.NewPropertyService()
	var propertyObj *propertyModel.Property
	propertyObj, err = service.Update(propertyId, req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSON(c, propertyObj)
}
