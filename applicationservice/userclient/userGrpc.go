// Package userclient is client library if you need to call the user Micro-service as a client.
// It provides client library and the data transformation service.
package userclient

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jfeng45/servicetmpl1/applicationservice/userclient/generatedclient"
	"github.com/jfeng45/servicetmpl1/domain/model"
	"github.com/pkg/errors"
)

// GrpcToUser converts from grpc User type to domain Model user type
func GrpcToUser(user *generatedclient.User) (*model.User, error) {
	if user == nil {
		return nil, nil
	}
	resultUser := model.User{}

	resultUser.Id = int(user.Id)
	resultUser.Name = user.Name
	resultUser.Department = user.Department
	created, err := ptypes.Timestamp(user.Created)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	resultUser.Created = created
	return &resultUser, nil
}

// UserToGrpc converts from domain Model User type to grpc user type
func UserToGrpc(user *model.User) (*generatedclient.User, error) {
	if user == nil {
		return nil, nil
	}
	resultUser := generatedclient.User{}
	resultUser.Id = int32(user.Id)
	resultUser.Name = user.Name
	resultUser.Department = user.Department
	created, err := ptypes.TimestampProto(user.Created)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	resultUser.Created = created
	return &resultUser, nil
}

// UserListToGrpc converts from array of domain Model User type to array of grpc user type
func UserListToGrpc(ul []model.User) ([]*generatedclient.User, error) {
	var gul []*generatedclient.User
	for _, user := range ul {
		gu, err := UserToGrpc(&user)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		gul = append(gul, gu)
	}
	return gul, nil
}
