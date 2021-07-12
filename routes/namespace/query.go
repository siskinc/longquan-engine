package namesapce

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
	namespaceService "github.com/siskinc/longquan-engine/service/namespace"
)

// QueryNamespace godoc
// @TAGS 命名空间
// @Summary 查询命名空间
// @Description 查询命名空间
// @Accept json
// @Produce json
// @Param message query namespaceService.QueryReq false "查询信息"
// @Success 200 {object} httpx.JSONResultPaged.{data=[]namespaceModel.Namespace}
// @Router /namespace [get]
func QueryNamespace(c *gin.Context) {
	var (
		req           = &namespaceService.QueryReq{}
		namespaceObjs []*namespaceModel.Namespace
		err           error
		total         int64
	)
	err = c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to QueryReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	service := namespaceService.NewService()
	namespaceObjs, total, err = service.Query(req)
	if err != nil {
		logrus.Errorf("query namespace have an err: %v, req: %+v", err, req)
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSONPaged(c, namespaceObjs, total)
}
