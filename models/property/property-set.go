package property

import "go.mongodb.org/mongo-driver/bson/primitive"

// PropertySet 属性集
// UniqueIndex NamespaceID + Code
type PropertySet struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`                    // 主键ID
	NamespaceID primitive.ObjectID `json:"namespace_id,omitempty" bson:"namespace_id,omitempty"` // 命名空间ID
	Code        string             `json:"code,omitempty"  bson:"code,omitempty"`                // 数据集CODE
	Description string             `json:"description,omitempty" bson:"description,omitempty"`   // 描述
}
