package property

import (
	"context"
	"fmt"

	"github.com/goools/tools/mongox"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *PropertyMongoRepository) UpdateOne(oid primitive.ObjectID, updater bson.M) (propertyObj *propertyModel.Property, err error) {
	if mongox.EmptyOid(oid) {
		err = fmt.Errorf("update one property have an err: object id is empty")
		return
	}
	filter := mongox.MakeQueryByID(oid)
	result := repo.collection.FindOneAndUpdate(context.Background(), filter, updater, mongox.MakeReturnAfter(nil))
	err = result.Err()
	if err != nil {
		return
	}
	propertyObj = &propertyModel.Property{}
	err = result.Decode(propertyObj)
	return
}
