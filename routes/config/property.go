package config

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/ginx/httpx"
	"github.com/siskinc/longquan-engine/constants/types"
)

// QueryProperty godoc
// @TAGS 场景管理
// @Summary 查询属性分类
// @Description 查询属性分类
// @Accept json
// @Produce json
// @Success 200 {object} httpx.JSONResult.{data=map[string]string}
// @Router /config/property-type-enum [get]
func PropertyTypeEnumData(c *gin.Context) {
	httpx.SetDefaultRespJSON(c, map[string]string{
		string(types.PropertyTypeEnumCustom): "自定义类型",
		string(types.PropertyTypeEnumSystem): "系统类型",
	})
}
