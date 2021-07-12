package namespace

import (
	"context"

	"github.com/goools/tools/mongox"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepository) Delete(oid primitive.ObjectID) (err error) {
	_, err = repo.collection.DeleteOne(context.Background(), mongox.MakeQueryByID(oid))
	return err
}
