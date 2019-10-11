package config

import (
	"github.com/sjuliper7/silhouette/common/models"
	"google.golang.org/grpc"

	"github.com/sjuliper7/silhouette/services/user-service/delivery"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
)

func (cf *Config) initService() {

	repo := repositories.NewMysqlRepository(cf.DB)
	usecase := usecase.NewUserUsecase(repo)

	svr := grpc.NewServer()
	userServer := delivery.NewUserServer(usecase)

	models.RegisterUsersServer(svr, userServer)
}
