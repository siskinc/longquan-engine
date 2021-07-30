package property

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *PropertySetService) checkExist(namespaceOid primitive.ObjectID, code string) (exist bool, err error) {
	filter := bson.D{
		bson.E{
			Key:   "namespace_id",
			Value: namespaceOid,
		},
		bson.E{
			Key:   "code",
			Value: code,
		},
	}
	var total int64
	_, total, err = service.propertySetMongoRepo.Qeury(filter, 0, 0, "")
	if err != nil {
		logrus.Errorf("check property set is exist have an err: %v, namespace id %s, code is %s", err, namespaceOid.Hex(), code)
		return
	}
	exist = total > 0
	return
}
