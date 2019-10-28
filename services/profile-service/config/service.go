package config

import (
	"github.com/sjuliper7/silhouette/common/config"
	"github.com/sjuliper7/silhouette/common/protocs"
	grpc2 "github.com/sjuliper7/silhouette/services/profile-service/delivery/grpc"
	"github.com/sjuliper7/silhouette/services/profile-service/repositories/mysql"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
)

func (cf *Config) initService() {
	repo := mysql.NewMysqlProfileRepository(cf.DB)
	profileUc := usecase.NewProfileUsecase(repo)

	svr := grpc.NewServer()
	profileServer := grpc2.NewProfileServer(profileUc)
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
