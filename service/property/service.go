package property

import (
	"github.com/siskinc/longquan-engine/global"
	propertyRepository "github.com/siskinc/longquan-engine/repository/property"
)

type PropertyService struct {
	mongoRepo *propertyRepository.PropertyMongoRepository
}

func NewService() *PropertyService {
	serviceObj := &PropertyService{
		mongoRepo: propertyRepository.NewPropertyMongoRepository(global.Config.MongoDbDriver),
	}
	return serviceObj
}
