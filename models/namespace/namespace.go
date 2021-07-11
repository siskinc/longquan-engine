package namespace

import "go.mongodb.org/mongo-driver/bson/primitive"

type Namespace struct {
	ID          primitive.ObjectID `json:"id,omitempty" yaml:"id,omitempty" bson:"_id,omitempty"`                           // 主键ID
	Name        string             `json:"name,omitempty" yaml:"name,omitempty" bson:"name,omitempty"`                      // 名称
	Description string             `json:"description,omitempty" yaml:"description,omitempty" bson:"description,omitempty"` // 描述
}

func NewNamespace(name, desc string) *Namespace {
	return &Namespace{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: desc,
	}
}
