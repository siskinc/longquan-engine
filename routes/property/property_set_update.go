package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// UpdatePropertySet godoc
// @TAGS 场景管理
// @Summary 修改属性集
// @Description 修改属性集
// @Accept json
// @Produce json
// @Param id path string true "属性集id" minlength(1)
// @Param message body propertyService.UpdatePropertySetReq true "属性集更新信息"
// @Success 200 {object} httpx.JSONResult.{data=propertyModel.PropertySet}
// @Router /property-set/{id} [patch]
func UpdatePropertySet(c *gin.Context) {
	propertyId := c.Param("id")
	req := &propertyService.UpdatePropertySetReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to UpdatePropertySetReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	service := propertyService.NewPropertySetService()
	var propertyObj *propertyModel.PropertySet
	propertyObj, err = service.Update(propertyId, req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSON(c, propertyObj)
}
