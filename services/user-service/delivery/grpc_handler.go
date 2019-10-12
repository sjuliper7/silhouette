package delivery

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sjuliper7/silhouette/common/models"
	"log"
)

func (us UserServer) List(context.Context, *empty.Empty) (*models.UserList, error) {
	var userList models.UserList
	var users []*models.User

	uu, err := us.usecase.GetAlluser()

	if err != nil {
		log.Println("Failed when call [usecase][GetAlluser]")
		return nil, err
	}

	for _, u := range uu {
		var user models.User
		user.Id = u.ID
		user.Name = u.Name
		user.LastName = u.LastName

		users = append(users, &user)
	}

	userList.List = users

	return &userList, nil
}
