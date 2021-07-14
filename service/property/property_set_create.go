package property

import (
	"fmt"

	"github.com/goools/tools/errorx"
	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	"github.com/siskinc/longquan-engine/global"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	namespaceRepository "github.com/siskinc/longquan-engine/repository/namespace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePropertySetReq struct {
	NamespaceID string `json:"namespace_id"` // 关联命名空间
	Code        string `json:"code"`         // 属性集code
	Name        string `json:"name"`         // 字段名
	Description string `json:"description"`  // 描述
}

func (service *PropertySetService) makePropertySetCodeExistFilter(namespaceOid primitive.ObjectID, code string) bson.D {
	return bson.D{
		bson.E{
			Key:   "namespace_id",
			Value: namespaceOid,
		},
		bson.E{
			Key:   "property_set_code",
			Value: code,
		},
	}
}

func (service *PropertySetService) Create(req *CreatePropertySetReq) (propertySetObj *propertyModel.PropertySet, err error) {
	var namespaceOid primitive.ObjectID
	namespaceOid, err = primitive.ObjectIDFromHex(req.NamespaceID)
	if err != nil {
		logrus.Errorf("namespace id is invalid when create property set, err: %v, id: %s", err, req.NamespaceID)
		err = error_code.ParameterInvalidNamespaceIDError
		return
	}
	// check namespace id
	namespaceRepo := namespaceRepository.NewMongoRepository(global.Config.MongoDbDriver)
	namespaceFilter := mongox.MakeQueryByID(namespaceOid)
	var total int64
	_, total, err = namespaceRepo.Query(namespaceFilter, 0, 0, "")
	if err != nil {
		logrus.Errorf("query namespace when create property set have an err: %v, namespace id: %s", err, req.NamespaceID)
		return
	}
	if total > 0 {
		logrus.Errorf("cannot find namespace with id: %s", req.NamespaceID)
		err = errorx.NewError(error_code.CustomForbiddenNotFoundNamespace, fmt.Errorf("not found namespace with id: %s", req.NamespaceID))
		return
	}

	// check property set code exist
	propertySetCodeExistFilter := service.makePropertySetCodeExistFilter(namespaceOid, req.Code)
	_, total, err = service.mongoRepo.Qeury(propertySetCodeExistFilter, 0, 0, "")
	if err != nil {
		logrus.Errorf("query property set when create property set have an err: %v, request: %+v", err, req)
		return
	}
	if total > 0 {
		logrus.Errorf("create property set have an err, code %s conflict", req.Code)
		err = errorx.NewError(error_code.CustomForbiddenConflictPropertySet, fmt.Errorf("property set code %s conflict", req.Code))
		return
	}

	// create property set
	propertySetObj = &propertyModel.PropertySet{
		ID:          primitive.NewObjectID(),
		NamespaceID: namespaceOid,
		Code:        req.Code,
		Description: req.Description,
	}
	err = service.mongoRepo.Create(propertySetObj)
	if err != nil {
		logrus.Errorf("create property set have an err: %v, property object: %+v", err, propertySetObj)
		return
	}
	return
}
