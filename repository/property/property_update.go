package property

import (
	"context"
	"fmt"

	"github.com/goools/tools/mongox"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (repo *PropertyMongoRepository) UpdateOne(oid primitive.ObjectID, updater bson.M) (PropertyObj *propertyModel.Property, err error) {
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
	PropertyObj = &propertyModel.Property{}
	err = result.Decode(PropertyObj)
	return
}
