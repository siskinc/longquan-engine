package property

import (
	"fmt"

	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	"github.com/siskinc/longquan-engine/constants/types"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	namespaceService "github.com/siskinc/longquan-engine/service/namespace"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePropertyReq struct {
	NamespaceID     string `json:"namespace_id,omitempty" bson:"namespace_id,omitempty"`           // 关联命名空间
	PropertySetCode string `json:"property_set_code,omitempty" bson:"property_set_code,omitempty"` // 关联属性集
	Name            string `json:"name,omitempty" bson:"name,omitempty"`                           // 字段名
	Class           string `json:"class,omitempty" bson:"class,omitempty"`                         // 字段类型 int, string, float, ${property_set_code}
	Type            string `json:"type,omitempty" bson:"type,omitempty"`                           // 字段分类 system or custome
	Description     string `json:"description,omitempty" bson:"description,omitempty"`             // 描述
}

func (service *PropertyService) CreateProperty(req *CreatePropertyReq) (propertyObj *propertyModel.Property, err error) {
	var namespaceOid primitive.ObjectID
	namespaceOid, err = primitive.ObjectIDFromHex(req.NamespaceID)
	if err != nil {
		logrus.Errorf("namespace id not is a valid object id, err: %v, request: %+v", err, req)
		err = error_code.NewParameterInvalid(fmt.Errorf("request namespace id is invalid"))
		return
	}

	var exist bool
	// check namespace exist
	namespaceServiceObj := namespaceService.NewService()
	exist, err = namespaceServiceObj.CheckExist(namespaceOid)
	if err != nil {
		return
	}
	if !exist {
		err = errorx.NewError(
			error_code.CustomForbiddenNotFoundNamespace,
			fmt.Errorf("not found namespace, namespace id: %s", req.NamespaceID),
		)
		return
	}

	// check property set code is valid
	propertySetService := NewPropertySetService()
	exist, err = propertySetService.checkExist(namespaceOid, req.PropertySetCode)
	if err != nil {
		return
	}
	if !exist {
		err = errorx.NewError(
			error_code.CustomForbiddenNotFoundPropertySet,
			fmt.Errorf("not found property set, namespace id: %s, code: %s", req.NamespaceID, req.PropertySetCode),
		)
		return
	}
	// check NamespaceID + PropertySetCode + Name exist
	exist, err = service.checkExist(namespaceOid, req.PropertySetCode, req.Name)
	if err != nil {
		return
	}
	if exist {
		err = errorx.NewError(
			error_code.CustomForbiddenConflictProperty,
			fmt.Errorf("property %s exist", req.Name),
		)
		return
	}

	propertyObj = &propertyModel.Property{
		ID:              primitive.NewObjectID(),
		NamespaceID:     namespaceOid,
		PropertySetCode: req.PropertySetCode,
		Name:            req.Name,
		Class:           types.PropertyTypeEnum(req.Class),
		Type:            req.Type,
		Description:     req.Description,
	}
	err = service.mongoRepo.Create(propertyObj)
	if err != nil {
		logrus.Errorf("create property have an err: %v, request: %+v", err, req)
		return
	}
	return
}
