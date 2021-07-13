package property

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
)

func (repo PropertyMongoRepository) Create(propertyObj *propertyModel.Property) (err error) {
	if propertyObj == nil {
		err = fmt.Errorf("property object is nil")
		return
	}
	_, err = repo.collection.InsertOne(context.Background(), propertyObj)
	if err != nil {
		logrus.Errorf("create namesapce to mongo have an err: %v, property object: %+v", err, propertyObj)
		return
	}
	return
}
