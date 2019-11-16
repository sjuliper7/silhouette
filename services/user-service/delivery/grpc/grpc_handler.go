package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/common/protocs"
)

func (us UserServer) List(context.Context, *empty.Empty) (*protocs.UserList, error) {
	var userList protocs.UserList
	var users []*protocs.User

	uu, err := us.usecase.GetAlluser()

	if err != nil {
		logrus.Println("Failed when call [usecase][GetAlluser]")
		return nil, err
	}

	for _, u := range uu {
		var user protocs.User
		user.ID = u.ID
		user.Email = u.Email
		user.Username = u.Username
		user.Name = u.Name
		user.Role = u.Role

		users = append(users, &user)
	}

	userList.List = users

	return &userList, nil
}
