package property

import (
	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryPropertySetReq struct {
	ID          *string `json:"id,omitempty"`           // 主键ID
	NamespaceID *string `json:"namespace_id,omitempty"` // 命名空间ID
	Code        *string `json:"code,omitempty"`         // 数据集CODE
	Description *string `json:"description,omitempty"`  // 描述
	PageIndex   int64   `form:"page_index"`             // 分页：查询第某页
	PageSize    int64   `form:"page_size"`              // 分页：每页数量
	SortedField *string `form:"sorted_field"`           // 排序字段，默认是以主键倒序，例如：-name,这个例子是以name为倒序排列
}

func (service *PropertySetService) QueryPropertySet(req *QueryPropertySetReq) (propertySetObjs []*propertyModel.PropertySet, total int64, err error) {
	filter := bson.D{}
	if req.ID != nil && *req.ID != "" {
		var oid primitive.ObjectID
		oid, err = primitive.ObjectIDFromHex(*req.ID)
		if err != nil {
			logrus.Errorf("id is invalid when query set property, request: %+v", req)
			err = error_code.ParameterInvalidIDError
			return
		}
		filter = mongox.MakeQueryByID(oid)
	} else {
		if req.NamespaceID != nil && *req.NamespaceID != "" {
			var namespaceOid primitive.ObjectID
			namespaceOid, err = primitive.ObjectIDFromHex(*req.NamespaceID)
			if err != nil {
				logrus.Errorf("namespace id is invalid, err: %v, namespace id: %s", err, *req.NamespaceID)
				err = error_code.ParameterInvalidNamespaceIDError
				return
			}
			filter = append(filter, bson.E{Key: "namespace_id", Value: namespaceOid})
		}
		if req.Code != nil && *req.Code != "" {
			filter = append(filter, bson.E{Key: "code", Value: *req.Code})
		}
		if req.Description != nil && *req.Description != "" {
			filter = append(filter, bson.E{Key: "description", Value: *req.Description})
		}
	}

	propertySetObjs, total, err = service.mongoRepo.Qeury(filter, req.PageIndex, req.PageSize, *req.SortedField)
	return
}
