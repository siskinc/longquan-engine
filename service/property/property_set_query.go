package property

import (
	"fmt"

	"github.com/goools/tools/ginx/paramsx"
	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryPropertySetReq struct {
	ID          *string `form:"id,omitempty"`           // 主键ID
	NamespaceID *string `form:"namespace_id,omitempty"` // 命名空间ID
	Code        *string `form:"code,omitempty"`         // 数据集CODE
	Description *string `form:"description,omitempty"`  // 描述
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

	sortedField := ""
	if req.SortedField != nil && *req.SortedField != "" {
		sortedField = *req.SortedField
	}
	pageIndex, pageSize := paramsx.ShiftPage(req.PageIndex, req.PageSize)
	propertySetObjs, total, err = service.propertySetMongoRepo.Qeury(filter, pageIndex, pageSize, sortedField)
	return
}

func (service *PropertySetService) QueryByID(id string) (propertySet *propertyModel.PropertySet, err error) {
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		err = error_code.ParameterInvalidIDError
		return
	}
	filter := mongox.MakeQueryByID(oid)
	var propertySetObjs []*propertyModel.PropertySet
	var total int64
	propertySetObjs, total, err = service.propertySetMongoRepo.Qeury(filter, 0, 0, "-_id")
	if err != nil {
		return
	}
	if total <= 0 {
		err = error_code.NewCustomForbiddenNotFoundPropertySetError(fmt.Errorf("cannot find property set by id, id: %s", id))
		return
	}
	propertySet = propertySetObjs[0]
	return
}

func (service *PropertySetService) QueryByNamespaceAndCode(namespaceOid primitive.ObjectID, code string) (propertySet *propertyModel.PropertySet, err error) {
	filter := bson.D{}
	filter = append(filter, bson.E{
		Key:   "namespace_id",
		Value: namespaceOid,
	})
	filter = append(filter, bson.E{
		Key:   "code",
		Value: code,
	})
	var propertySetObjs []*propertyModel.PropertySet
	var total int64
	propertySetObjs, total, err = service.propertySetMongoRepo.Qeury(filter, 0, 0, "-_id")
	if err != nil {
		return
	}
	if total <= 0 {
		err = error_code.NewCustomForbiddenNotFoundPropertySetError(
			fmt.Errorf("cannot find property set by namespace_id and code, namespace_id: %s, code: %s", namespaceOid.String(), code),
		)
		return
	}
	propertySet = propertySetObjs[0]
	return
}
