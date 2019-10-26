package config

import (
	"github.com/sjuliper7/silhouette/common/config"
	"github.com/sjuliper7/silhouette/common/protocs"
	"github.com/sjuliper7/silhouette/services/profile-service/delivery"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/sjuliper7/silhouette/services/profile-service/repositories"
	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
)

func (cf *Config) initService() {
	repo := repositories.NewMysqlRepository(cf.DB)
	usecase := usecase.NewProfileUsecase(repo)

	svr := grpc.NewServer()
	profileServer := delivery.NewProfileServer(usecase)
	//
	protocs.RegisterProfilesServer(svr, profileServer)
	log.Println("Starting RPC server at", config.SERVICE_PROFILE_PORT)

	//next running the to http
	net, err := net.Listen("tcp", config.SERVICE_PROFILE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_PROFILE_PORT, err)
	}

	log.Fatalln(svr.Serve(net))
}
