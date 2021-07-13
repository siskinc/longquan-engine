package property

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *PropertyService) DeleteOne(id string) (err error) {
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Errorf("property id is invalid when delete ,err: %v, id: %s", err, id)
		err = error_code.ParameterInvalidIDError
		return
	}
	err = service.mongoRepo.Delete(oid)
	if err != nil {
		logrus.Errorf("delete property have an err: %v, id: %s", err, id)
		return
	}
	return
}
