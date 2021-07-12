package namesapce

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	namespaceService "github.com/siskinc/longquan-engine/service/namespace"
)

// UpdateNamespace godoc
// @TAGS 命名空间
// @Summary 修改命名空间
// @Description 修改命名空间
// @Accept json
// @Produce json
// @Param id path string true "命名空间id" minlength(1)
// @Param message body namespaceService.UpdateOneReq true "命名空间更新信息"
// @Success 200 {object} httpx.JSONResult
// @Router /namespace/{id} [patch]
func UpdateNamespace(c *gin.Context) {
	namespaceId := c.Param("id")
	req := &namespaceService.UpdateOneReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to UpdateOneReq, err: %v", err)
		httpx.SetRespErr(c, error_code.NewParameterInvalid(err))
		return
	}
	service := namespaceService.NewService()
	namespaceObj, err := service.UpdateOne(namespaceId, req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetDefaultRespJSON(c, namespaceObj)
}
