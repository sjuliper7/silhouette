package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/models"
)

func (us UserServer) List(context.Context, *empty.Empty) (*models.UserList, error) {
	var userList models.UserList
	var users []*models.User

	uu, err := us.usecase.GetAll()

	if err != nil {
		logrus.Println("Failed when call [usecase][GetAlluser] %v", err)
		return nil, err
	}

	for _, u := range uu {
		var user models.User
		user.ID = uint64(u.ID)
		user.Email = u.Email
		user.Username = u.Username
		user.Name = u.Name
		user.Role = u.Role

		users = append(users, &user)
	}

	userList.List = users

	return &userList, nil
}
