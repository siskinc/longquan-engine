package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// CreatePropertySet godoc
// @TAGS 场景管理
// @Summary 创建属性集
// @Description 创建属性集
// @Accept json
// @Produce json
// @Param message body propertyService.CreatePropertySetReq true "属性集基本信息"
// @Success 200 {object} httpx.JSONResult.{data=propertyModel.PropertySet} ""
// @Router /property-set [post]
func CreatePropertySet(c *gin.Context) {
	req := &propertyService.CreatePropertySetReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to CreatePropertySetReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	var propertyObj *propertyModel.PropertySet
	service := propertyService.NewPropertySetService()
	propertyObj, err = service.Create(req)
	if err != nil {
		logrus.Errorf("create property set have an err: %v", err)
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSON(c, propertyObj)
	logrus.Infof("create property set success, property set object: %+v", propertyObj)
}
