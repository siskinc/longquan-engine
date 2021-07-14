package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// CreateProperty godoc
// @TAGS 场景管理
// @Summary 创建属性
// @Description 创建属性
// @Accept json
// @Produce json
// @Param message body propertyService.CreatePropertyReq true "属性基本信息"
// @Success 200 {object} httpx.JSONResult.{data=propertyModel.Property} ""
// @Router /property [post]
func CreateProperty(c *gin.Context) {
	req := &propertyService.CreatePropertyReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to CreatePropertyReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	var propertyObj *propertyModel.Property
	service := propertyService.NewPropertyService()
	propertyObj, err = service.CreateProperty(req)
	if err != nil {
		logrus.Errorf("create property have an err: %v", err)
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSON(c, propertyObj)
	logrus.Infof("create property success, property object: %+v", propertyObj)
}
