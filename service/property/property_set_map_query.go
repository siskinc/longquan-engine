package property

import (
	"fmt"

	"github.com/goools/tools/errorx"
	"github.com/siskinc/longquan-engine/constants/error_code"
	"github.com/siskinc/longquan-engine/constants/types"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *PropertySetMapService) Query(id string) (propertyMap *propertyModel.PropertyMap, err error) {
	queryPSReq := &QueryPropertySetReq{
		ID: &id,
	}
	var propertySetObjs []*propertyModel.PropertySet
	var total int64
	propertySetObjs, total, err = service.propertySetService.QueryPropertySet(queryPSReq)
	if err != nil {
		return
	}
	if total <= 0 {
		err = errorx.NewError(
			error_code.CustomForbiddenNotFoundPropertySet,
			fmt.Errorf("not found property set, property set id: %s", id),
		)
		return
	}
	propertySetObj := propertySetObjs[0]
	propertyMap, err = service.buildPropertySetMap(propertySetObj)
	return
}

func (service *PropertySetMapService) buildPropertySetMap(propertySetObj *propertyModel.PropertySet) (
	propertyMap *propertyModel.PropertyMap, err error) {
	propertyMap = &propertyModel.PropertyMap{}
	propertySetCode := propertySetObj.Code
	namespaceID := propertySetObj.NamespaceID
	var propertyObjs []*propertyModel.Property
	filter := bson.D{
		bson.E{
			Key:   "namespace_id",
			Value: namespaceID,
		},
		bson.E{
			Key:   "property_set_code",
			Value: propertySetCode,
		},
	}
	propertyObjs, _, err = service.propertyService.mongoRepo.Query(filter, 0, 0, "")
	if err != nil {
		err = fmt.Errorf("query property have an err: %v, namespace_id: %v, property_set_code: %s", err, namespaceID, propertySetCode)
		return
	}
	var (
		propertyMapSub *propertyModel.PropertyMap
		propertyObj    *propertyModel.Property
	)
	for i := range propertyObjs {
		propertyObj = propertyObjs[i]
		if propertyObj.Type == types.PropertyTypeEnumCustom {
			propertyMapSub, err = service.buildPropertySetMap(&propertyModel.PropertySet{
				NamespaceID: namespaceID,
				Code:        propertyObj.PropertySetCode,
			})
			if err != nil {
				return
			}
		} else {
			propertyMapSub = &propertyModel.PropertyMap{
				Property: *propertyObj,
			}
		}
		propertyMap.Children = append(propertyMap.Children, propertyMapSub)
	}
	return
}
