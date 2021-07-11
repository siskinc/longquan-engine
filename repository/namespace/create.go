package namespace

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
)

// Create create a new namespace object to mongodb
func (repo *MongoRepository) Create(namespaceObj *namespaceModel.Namespace) (err error) {
	if namespaceObj == nil {
		err = fmt.Errorf("namespace object is nil")
		return
	}
	_, err = repo.collection.InsertOne(context.Background(), namespaceObj)
	if err != nil {
		logrus.Errorf("create namesapce to mongo have an err: %v, namespace object: %+v", err, namespaceObj)
		return
	}
	return
}
