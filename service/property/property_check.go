package property

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *PropertyService) checkExist(namespaceOid primitive.ObjectID, propertySetCode, name string) (exist bool, err error) {
	filter := bson.D{
		bson.E{
			Key:   "namespace_id",
			Value: namespaceOid,
		},
		bson.E{
			Key:   "property_set_code",
			Value: propertySetCode,
		},
		bson.E{
			Key:   "name",
			Value: name,
		},
	}
	var total int64
	_, total, err = service.mongoRepo.Query(filter, 0, 0, "")
	if err != nil {
		logrus.Errorf("check property exist have an err: %v, namespace id: %s, property set code: %s, property name: %s",
			err, namespaceOid.Hex(), propertySetCode, name)
		return
	}
	exist = total > 0
	return
}
