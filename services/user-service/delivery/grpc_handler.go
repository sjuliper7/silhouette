package delivery

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sjuliper7/silhouette/common/models"
)

func (us UserServer) List(context.Context, *empty.Empty) (*models.UserList, error) {
	var users *models.UserList

	users.List = make([]*models.User, 0)

	uu := us.usecase.GetAlluser()

	for _, u := range uu {
		var user *models.User
		user.Id = u.ID
		user.Name = u.Name
		user.LastName = u.LastName

		users.List = append(users.List, user)
	}

	return users, nil
}
