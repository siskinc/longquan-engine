package property

import (
	"github.com/siskinc/longquan-engine/constants/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Property 属性
// Unique Index: NamespaceID PropertySetCode Name
type Property struct {
	ID              primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`                              // 主键ID
	NamespaceID     primitive.ObjectID     `json:"namespace_id,omitempty" bson:"namespace_id,omitempty"`           // 关联命名空间
	PropertySetCode string                 `json:"property_set_code,omitempty" bson:"property_set_code,omitempty"` // 关联属性集
	Name            string                 `json:"name,omitempty" bson:"name,omitempty"`                           // 字段名
	Class           string                 `json:"class,omitempty" bson:"class,omitempty"`                         // 字段类型 int, string, float, ${property_set_code}
	Type            types.PropertyTypeEnum `json:"type,omitempty" bson:"type,omitempty"`                           // 字段分类 system or custome
	Description     string                 `json:"description,omitempty" bson:"description,omitempty"`             // 描述
	PropertySet     *PropertySet           `json:"property_set" bson:"-"`                                          // 如果属性的类型是自定义类型, 那么会将对应的property-set取出
}
