package namespace

import (
	"github.com/siskinc/longquan-engine/constants/types"
	"github.com/siskinc/longquan-engine/driver/mongodb"
	"github.com/siskinc/longquan-engine/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(driver *mongodb.MongoDbDriver) *MongoRepository {
	repo := &MongoRepository{}
	repo.collection = driver.Collection(global.Config.MongoDbDriver.DatabaseName, types.CollectionNamespace)
	return repo
}
