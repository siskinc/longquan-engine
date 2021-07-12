package namespace

import (
	"context"
	"fmt"

	"github.com/goools/tools/mongox"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (repo *MongoRepository) UpdateOne(oid primitive.ObjectID, updater bson.M) (namespaceObj *namespaceModel.Namespace, err error) {
	if mongox.EmptyOid(oid) {
		err = fmt.Errorf("update one namespace have an err: object id is empty")
		return
	}
	filter := mongox.MakeQueryByID(oid)
	result := repo.collection.FindOneAndUpdate(context.Background(), filter, updater, mongox.MakeReturnAfter(nil))
	err = result.Err()
	if err != nil {
		return
	}
	namespaceObj = &namespaceModel.Namespace{}
	err = result.Decode(namespaceObj)
	return
}
