package property

import (
	"github.com/goools/tools/ginx/paramsx"
	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryPropertyReq struct {
	ID              *string `form:"id"`                // 主键id
	NamespaceID     *string `form:"namespace_id"`      // 命名空间ID
	PropertySetCode *string `form:"property_set_code"` // 属性集code
	Name            *string `form:"name"`              // 名称(模糊查询)
	Description     *string `form:"description"`       // 描述(模糊查询)
	PageIndex       int64   `form:"page_index"`        // 分页：查询第某页
	PageSize        int64   `form:"page_size"`         // 分页：每页数量
	SortedField     *string `form:"sorted_field"`      // 排序字段，默认是以主键倒序，例如：-name,这个例子是以name为倒序排列
}

func (service *PropertyService) Query(req *QueryPropertyReq) (propertyObjs []*propertyModel.Property, total int64, err error) {
	var filter = bson.D{}
	if req.ID != nil && *req.ID != "" {
		var oid primitive.ObjectID
		oid, err = primitive.ObjectIDFromHex(*req.ID)
		if err != nil {
			logrus.Errorf("id is invalid when query property, request: %+v", req)
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
		if req.PropertySetCode != nil && *req.PropertySetCode != "" {
			filter = append(filter, bson.E{Key: "property_set_code", Value: *req.PropertySetCode})
		}
		if req.Name != nil && *req.Name != "" {
			filter = append(filter, bson.E{Key: "name", Value: *req.Name})
		}
		if req.Description != nil && *req.Description != "" {
			filter = append(filter, bson.E{Key: "description", Value: *req.Description})
		}
	}
	pageIndex, pageSize := paramsx.ShiftPage(req.PageIndex, req.PageSize)
	propertyObjs, total, err = service.mongoRepo.Query(filter, pageIndex, pageSize, *req.SortedField)
	if err != nil {
		logrus.Errorf("query property is err: %v, request: %+v", err, req)
		return
	}
	return
}
