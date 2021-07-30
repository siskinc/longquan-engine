package property

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/types"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *PropertySetService) QueryPropertyByPorpertySet(propertySet *propertyModel.PropertySet) (propertyObjs []*propertyModel.Property, err error) {
	namespaceOid := propertySet.NamespaceID
	propertySetCode := propertySet.Code

	filter := bson.D{}
	filter = append(filter, bson.E{
		Key:   "namespace_id",
		Value: namespaceOid,
	})
	filter = append(filter, bson.E{
		Key:   "property_set_code",
		Value: propertySetCode,
	})
	propertyObjs, _, err = service.propertyMongoRepo.Query(filter, 0, 0, "-_id")
	return
}

func (service *PropertySetService) FullPropertySet(propertySet *propertyModel.PropertySet) (err error) {
	fulledPropertySetMap := map[string]*propertyModel.PropertySet{
		propertySet.Code: propertySet,
	}
	var waitFullPropertySetQueue []*propertyModel.PropertySet
	waitFullPropertySetQueue = append(waitFullPropertySetQueue, propertySet)
	for len(waitFullPropertySetQueue) > 0 {
		waitFullPropertySet := waitFullPropertySetQueue[0]
		waitFullPropertySetQueue = waitFullPropertySetQueue[1:]
		if waitFullPropertySet.ID.IsZero() {
			var tmpPropertySet *propertyModel.PropertySet
			tmpPropertySet, err = service.QueryByNamespaceAndCode(waitFullPropertySet.NamespaceID, waitFullPropertySet.Code)
			if err != nil {
				return
			}
			waitFullPropertySet.ID = tmpPropertySet.ID
			waitFullPropertySet.Description = tmpPropertySet.Description
		}
		waitFullPropertySet.PropertyList, err = service.QueryPropertyByPorpertySet(propertySet)
		if err != nil {
			return
		}
		for _, property := range waitFullPropertySet.PropertyList {
			if property.Type == types.PropertyTypeEnumSystem {
				continue
			}
			propertySetCode := property.PropertySetCode
			_propertySet, exist := fulledPropertySetMap[propertySetCode]
			if exist {
				_propertySet = &propertyModel.PropertySet{
					NamespaceID: property.NamespaceID,
					Code:        propertySetCode,
				}
				fulledPropertySetMap[propertySetCode] = _propertySet
				waitFullPropertySetQueue = append(waitFullPropertySetQueue, _propertySet)
			}
			property.PropertySet = _propertySet
		}
	}
	return
}

func (service *PropertySetService) QueryPropertySetMap(id string) (propertySet *propertyModel.PropertySet, err error) {
	propertySet, err = service.QueryByID(id)
	if err != nil {
		logrus.Errorf("query property by id have an err: %+v, id: %s", err, id)
		return
	}
	service.FullPropertySet(propertySet)
	return
}
