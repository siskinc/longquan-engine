package property

import "github.com/siskinc/longquan-engine/constants/types"

// Property 属性
type Property struct {
	Name        string                 `json:"name,omitempty" bson:"name,omitempty"`               // 字段名
	Class       types.PropertyTypeEnum `json:"class,omitempty" bson:"class,omitempty"`             // 字段类型 int, string, float, ${property_set_code}
	Type        string                 `json:"type,omitempty" bson:"type,omitempty"`               // 字段分类 system or custome
	Description string                 `json:"description,omitempty" bson:"description,omitempty"` // 描述
}
