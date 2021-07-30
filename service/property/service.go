package property

import (
	"github.com/siskinc/longquan-engine/global"
	propertyRepository "github.com/siskinc/longquan-engine/repository/property"
)

type PropertyService struct {
	mongoRepo *propertyRepository.PropertyMongoRepository
}

func NewPropertyService() *PropertyService {
	serviceObj := &PropertyService{
		mongoRepo: propertyRepository.NewPropertyMongoRepository(global.Config.MongoDbDriver),
	}
	return serviceObj
}

type PropertySetService struct {
	propertySetMongoRepo *propertyRepository.PropertySetMongoRepository
	propertyMongoRepo    *propertyRepository.PropertyMongoRepository
}

func NewPropertySetService() *PropertySetService {
	serviceObj := &PropertySetService{
		propertySetMongoRepo: propertyRepository.NewPropertySetMongoRepository(global.Config.MongoDbDriver),
		propertyMongoRepo:    propertyRepository.NewPropertyMongoRepository(global.Config.MongoDbDriver),
	}
	return serviceObj
}
