package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// QueryPropertySet godoc
// @TAGS 场景管理
// @Summary 查询属性集
// @Description 查询属性集
// @Accept json
// @Produce json
// @Param message query propertyService.QueryPropertySetReq false "查询信息"
// @Success 200 {object} httpx.JSONResultPaged.{data=[]propertyModel.PropertySet}
// @Router /namespace-set [get]
func QueryPropertySet(c *gin.Context) {
	var (
		req          = &propertyService.QueryPropertySetReq{}
		propertyObjs []*propertyModel.PropertySet
		err          error
		total        int64
	)
	err = c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to QueryPropertySetReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	service := propertyService.NewPropertySetService()
	propertyObjs, total, err = service.QueryPropertySet(req)
	if err != nil {
		logrus.Errorf("query property set have an err: %v, req: %+v", err, req)
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSONPaged(c, propertyObjs, total)
}
