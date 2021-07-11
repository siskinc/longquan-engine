package namespace

import (
	"github.com/siskinc/longquan-engine/global"
	namespaceRepository "github.com/siskinc/longquan-engine/repository/namespace"
)

type Service struct {
	mongoRepo *namespaceRepository.MongoRepository
}

func NewService() *Service {
	serviceObj := &Service{
		mongoRepo: namespaceRepository.NewMongoRepository(global.Config.MongoDbDriver),
	}
	return serviceObj
}
