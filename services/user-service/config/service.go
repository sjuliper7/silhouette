package config

import (
	"github.com/sjuliper7/silhouette/common/config"
	"github.com/sjuliper7/silhouette/common/models"
	"google.golang.org/grpc"
	"log"

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
	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	//next running the to http
}
