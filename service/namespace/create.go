package namespace

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/global"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
	namespaceRepository "github.com/siskinc/longquan-engine/repository/namespace"
)

type CreateNamespaceReq struct {
	Name        string `json:"name" binding:"required" minlength:"1" example:"longquan-sword"`
	Description string `json:"description,omitempty" example:"rule-engine"`
}

func (service *Service) CraeteNamespace(req *CreateNamespaceReq) (namespaceObj *namespaceModel.Namespace, err error) {
	namespaceObj = namespaceModel.NewNamespace(req.Name, req.Description)
	repo := namespaceRepository.NewMongoRepository(global.Config.MongoDbDriver)
	err = repo.Create(namespaceObj)
	if err != nil {
		logrus.Errorf("create new namespace have an err: %v, namespace object: %+v", err, namespaceObj)
		return
	}
	logrus.Infof("create new namespace success, namespace object: %+v", namespaceObj)
	return
}
