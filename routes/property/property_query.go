package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// QueryProperty godoc
// @TAGS 场景管理
// @Summary 查询属性
// @Description 查询属性
// @Accept json
// @Produce json
// @Param message query propertyService.QueryPropertyReq false "查询信息"
// @Success 200 {object} httpx.JSONResultPaged.{data=[]propertyModel.Property}
// @Router /namespace [get]
func QueryProperty(c *gin.Context) {
	var (
		req          = &propertyService.QueryPropertyReq{}
		propertyObjs []*propertyModel.Property
		err          error
		total        int64
	)
	err = c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to QueryPropertyReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	service := propertyService.NewPropertyService()
	propertyObjs, total, err = service.Query(req)
	if err != nil {
		logrus.Errorf("query property have an err: %v, req: %+v", err, req)
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSONPaged(c, propertyObjs, total)
}
