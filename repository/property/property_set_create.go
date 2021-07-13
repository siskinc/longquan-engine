package property

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
)

func (repo *PropertySetMongoRepository) Create(propertySetObj *propertyModel.PropertySet) (err error) {
	if propertySetObj == nil {
		err = fmt.Errorf("property set object is nil")
		return
	}
	_, err = repo.collection.InsertOne(context.Background(), propertySetObj)
	if err != nil {
		logrus.Errorf("create namesapce to mongo have an err: %v, property set object: %+v", err, propertySetObj)
		return
	}
	return
}
