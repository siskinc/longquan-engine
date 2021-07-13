package property

import (
	"context"

	"github.com/goools/tools/mongox"
	"github.com/sirupsen/logrus"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *PropertySetMongoRepository) Qeury(filter bson.D, pageIndex, pageSize int64, sortedField string) (propertySetObjs []*propertyModel.PropertySet, total int64, err error) {
	opt := mongox.MakeFindPageOpt(nil, pageIndex, pageSize)
	opt = mongox.MakeSortedFieldOpt(opt, sortedField)
	total, err = repo.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		logrus.Errorf("query property set count have an err: %v, filter: %+v, sorted field: %s", err, filter, sortedField)
		return
	}
	var cursor *mongo.Cursor
	cursor, err = repo.collection.Find(context.Background(), filter, opt)
	if err != nil {
		logrus.Errorf("query property set have an err: %v, filter: %+v, sorted field: %s, page index: %d, page size: %d", err, filter, sortedField, pageIndex, pageSize)
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		propertySetObj := &propertyModel.PropertySet{}
		err = cursor.Decode(propertySetObj)
		if err != nil {
			logrus.Errorf("decode data to property set object have an err: %v", err)
			return
		}
		propertySetObjs = append(propertySetObjs, propertySetObj)
	}
	return
}
