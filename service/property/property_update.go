package property

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type UpdatePropertyReq struct {
	Description string `json:"description,omitempty" bson:"description,omitempty"` // 描述
}

func (service *PropertyService) Update(id string, req *UpdatePropertyReq) (propertyObj *propertyModel.Property, err error) {
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Errorf("property id is invalid when update, err: %v, id: %s", err, id)
		err = error_code.ParameterInvalidIDError
		return
	}
	propertyObj, err = service.mongoRepo.UpdateOne(oid, bson.M{"$set": bson.M{"description": req.Description}})
	if err != nil {
		logrus.Errorf("update property have an err: %v, id: %s, request: %+v", err, id, req)
		return
	}
	return
}
