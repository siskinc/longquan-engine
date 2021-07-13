package property

import (
	"github.com/siskinc/longquan-engine/constants/types"
	"github.com/siskinc/longquan-engine/driver/mongodb"
	"github.com/siskinc/longquan-engine/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type PropertyMongoRepository struct {
	collection *mongo.Collection
}

func NewPropertyMongoRepository(driver *mongodb.MongoDbDriver) *PropertyMongoRepository {
	repo := &PropertyMongoRepository{}
	repo.collection = driver.Collection(global.Config.MongoDbDriver.DatabaseName, types.CollectionProperty)
	return repo
}

type PropertySetMongoRepository struct {
	collection *mongo.Collection
}

func NewPropertySetMongoRepository(driver *mongodb.MongoDbDriver) *PropertySetMongoRepository {
	repo := &PropertySetMongoRepository{}
	repo.collection = driver.Collection(global.Config.MongoDbDriver.DatabaseName, types.CollectionPropertySet)
	return repo
}
