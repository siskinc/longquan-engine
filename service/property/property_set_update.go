package property

import (
	"github.com/sirupsen/logrus"
	"github.com/siskinc/longquan-engine/constants/error_code"
	propertyModel "github.com/siskinc/longquan-engine/models/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdatePropertySetReq struct {
	Description string `json:"description,omitempty"` // 描述
}

func (service *PropertySetService) Update(id string, req *UpdatePropertySetReq) (propertySetObj *propertyModel.PropertySet, err error) {
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Errorf("id is invalid when update property set, err: %v, id: %s, req: %+v", err, id, req)
		err = error_code.ParameterInvalidIDError
		return
	}

	updater := bson.M{
		"$set": bson.M{
			"description": req.Description,
		},
	}
	propertySetObj, err = service.mongoRepo.UpdateOne(oid, updater)
	if err != nil {
		logrus.Errorf("update property set have an err: %v, id: %s, req: %+v, updater: %+v", err, id, req, updater)
		return
	}
	return
}
