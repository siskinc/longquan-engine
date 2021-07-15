package namesapce

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/sirupsen/logrus"
	namespaceService "github.com/siskinc/longquan-engine/service/namespace"
)

// DeleteNamespace godoc
// @TAGS 命名空间
// @Summary 删除命名空间
// @Description 删除命名空间
// @Accept json
// @Produce json
// @Param id path string true "命名空间id" minlength(1)
// @Success 200 {object} httpx.JSONResult
// @Router /namespace/{id} [delete]
func DeleteNamespace(c *gin.Context) {
	namespaceId := c.Param("id")
	service := namespaceService.NewService()
	err := service.DeleteOne(namespaceId)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	logrus.Errorf("delete %s namespace success", namespaceId)
	httpx.SetDefaultRespJSON(c, nil)
}
