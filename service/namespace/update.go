package namespace

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	namespaceModel "github.com/siskinc/longquan-engine/models/namespace"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type UpdateOneReq struct {
	Name        *string `json:"name,omitempty" example:"longquan-sword"`
	Description *string `json:"description,omitempty" example:"rule-engine"`
}

func (service *Service) makeUpdater(req *UpdateOneReq) (result bson.M) {
	result = bson.M{}
	setter := bson.M{}
	if req.Name != nil && *req.Name != "" {
		setter["name"] = *req.Name
	}
	if req.Description != nil && *req.Description != "" {
		setter["description"] = *req.Description
	}
	result["$set"] = setter
	return
}

func (service *Service) UpdateOne(id string, req *UpdateOneReq) (namespaceObj *namespaceModel.Namespace, err error) {
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Errorf("request id parameter is invalid, id is %s, err: %v", id, err)
		err = error_code.ParameterInvalidIDError
		return
	}
	updater := service.makeUpdater(req)
	namespaceObj, err = service.mongoRepo.UpdateOne(oid, updater)
	if err != nil {
		logrus.Errorf("update namespace have an err: %v, id: %s, req: %+v", err, id, req)
		return
	}
	logrus.Infof("update namespace success, id: %s, req: %+v", id, req)
	return
}
