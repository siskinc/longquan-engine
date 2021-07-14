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
	mongoRepo *propertyRepository.PropertySetMongoRepository
}

func NewPropertySetService() *PropertySetService {
	serviceObj := &PropertySetService{
		mongoRepo: propertyRepository.NewPropertySetMongoRepository(global.Config.MongoDbDriver),
	}
	return serviceObj
}
