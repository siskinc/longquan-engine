package namespace

import (
	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryReq struct {
	ID          *string `form:"id"`           // 主键id
	Name        *string `form:"name"`         // 名称(模糊查询)
	Description *string `form:"description"`  // 描述(模糊查询)
	PageIndex   int64   `form:"page_index"`   // 分页：查询第某页
	PageSize    int64   `form:"page_size"`    // 分页：每页数量
	SortedField *string `form:"sorted_field"` // 排序字段，默认是以主键倒序，例如：-name,这个例子是以name为倒序排列
}

func (service *Service) Query(req *QueryReq) (namespaceObjs []*namespaceModel.Namespace, total int64, err error) {
	var filter = bson.D{}
	if req.ID != nil && *req.ID != "" {
		var oid primitive.ObjectID
		oid, err = primitive.ObjectIDFromHex(*req.ID)
		if err != nil {
			logrus.Errorf("query namespace have an err: %v, id is invalid: %s", err, *req.ID)
			err = error_code.ParameterInvalidIDError
			return
		}
		filter = mongox.MakeQueryByID(oid)
	} else {
		if req.Name != nil && *req.Name != "" {
			filter = append(filter, mongox.MakeRegexFilter("name", *req.Name))
		}
		if req.Description != nil && *req.Description != "" {
			filter = append(filter, mongox.MakeRegexFilter("description", *req.Description))
		}
	}
	sortedField := ""
	if req.SortedField != nil && *req.SortedField != "" {
		sortedField = *req.SortedField
	}
	namespaceObjs, total, err = service.mongoRepo.Query(filter, req.PageIndex, req.PageSize, sortedField)
	return
}
