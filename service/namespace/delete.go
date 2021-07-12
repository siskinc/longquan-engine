package namespace

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *Service) DeleteOne(id string) (err error) {
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Errorf("request id parameter is invalid, id is %s, err: %v", id, err)
		err = error_code.ParameterInvalidIDError
		return
	}
	err = service.mongoRepo.Delete(oid)
	if err != nil {
		logrus.Errorf("delete namespace have an err: %v, id: %s", err, id)
		return
	}
	return
}
