package property

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	propertyService "github.com/siskinc/longquan-engine/service/property"
)

// PropertySerMapQuery godoc
// @TAGS 场景管理
// @Summary 查询属性地图
// @Description 查询属性地图
// @Accept json
// @Produce json
// @Param id query string true "属性集id" minlength(1)
// @Success 200 {object} httpx.JSONResultPaged.{data=[]propertyModel.PropertyMap}
// @Router /property-set-map [get]
func PropertySerMapQuery(c *gin.Context) {
	id := c.Query("id")
	service := propertyService.NewPropertySetMapService()
	var propertyMap *propertyModel.PropertyMap
	var err error
	propertyMap, err = service.Query(id)
	if err != nil {
		logrus.Errorf("query property map have an err: %v, id is %s", err, id)
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSON(c, propertyMap)
}
