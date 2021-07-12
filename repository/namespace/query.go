package namespace

import (
	"context"

	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongoRepository) Query(filter bson.D, pageIndex, pageSize int64, sortedField string) (namespaceObjs []*namespaceModel.Namespace, total int64, err error) {
	opt := mongox.MakeFindPageOpt(nil, pageIndex, pageSize)
	opt = mongox.MakeSortedFieldOpt(opt, sortedField)
	total, err = repo.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		logrus.Errorf("query namespace count have an err: %v, filter: %+v, sorted field: %s", err, filter, sortedField)
		return
	}
	var cursor *mongo.Cursor
	cursor, err = repo.collection.Find(context.Background(), filter, opt)
	if err != nil {
		logrus.Errorf("query namespace have an err: %v, filter: %+v, sorted field: %s, page index: %d, page size: %d", err, filter, sortedField, pageIndex, pageSize)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		namespaceObj := &namespaceModel.Namespace{}
		err = cursor.Decode(namespaceObj)
		if err != nil {
			logrus.Errorf("decode data to namespace object have an err: %v", err)
			return
		}
		namespaceObjs = append(namespaceObjs, namespaceObj)
	}
	return
}
