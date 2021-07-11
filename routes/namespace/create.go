package namesapce

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
	namespaceService "github.com/siskinc/longquan-engine/service/namespace"
)

// CreateNamespace godoc
// @TAGS 命名空间
// @Summary 创建命名空间
// @Description 创建命名空间
// @Accept json
// @Produce json
// @Param message body namespaceService.CreateNamespaceReq true "命名空间属性"
// @Success 200 {object} httpx.JSONResult.{data=namespaceModel.Namespace} ""
// @Router /namespace [post]
func CreateNamespace(c *gin.Context) {
	var (
		namespaceObj *namespaceModel.Namespace
		err          error
	)
	req := &namespaceService.CreateNamespaceReq{}
	err = c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to CreateNamespaceReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	serviceObj := namespaceService.NewService()
	namespaceObj, err = serviceObj.CraeteNamespace(req)
	if err != nil {
		logrus.Errorf("create namespace have an err: %v, namesapce request: %+v", err, req)
		httpx.SetRespErr(c, err)
		return
	}
	logrus.Infof("create namespace success, namespace: %+v", namespaceObj)
	httpx.SetDefaultRespJSON(c, namespaceObj)
}
