package grpc

import (
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
)

// UserServer struct is  a struct to implement generated interface from proto
type UserService struct {
	usecase usecase.UserUsecase
}

//NewUserServer is function to implement usecase interface
func NewUserServer(uc usecase.UserUsecase) UserService {
	return UserService{uc}
}
