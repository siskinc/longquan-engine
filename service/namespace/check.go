package namespace

import (
	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *Service) CheckExist(namespaceOid primitive.ObjectID) (exist bool, err error) {
	filter := mongox.MakeQueryByID(namespaceOid)
	var total int64
	_, total, err = service.mongoRepo.Query(filter, 0, 0, "")
	if err != nil {
		logrus.Errorf("find namespace by id have an err: %v, id: %s", err, namespaceOid.Hex())
		return
	}
	exist = total > 0
	return
}
