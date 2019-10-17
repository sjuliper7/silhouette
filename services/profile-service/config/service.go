package config

import (
	"github.com/sjuliper7/silhouette/common/protocs"
	"google.golang.org/grpc"
	"log"
	"net"
)

func (cf *Config) initService() {
	//repo := repositories.NewMysqlRepository(cg.DB)
	//usecase := usecase.NewUserUsecase(repo)
	//
	//svr := grpc.NewServer()
	//userServer := delivery.NewUserServer(usecase)
	//
	//protocs.RegisterUsersServer(svr, userServer)
	//log.Println("Starting RPC server at", config.SERVICE_USER_PORT)
	//
	////next running the to http
	//net, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	//if err != nil {
	//	log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	//}
	//
	//log.Fatalln(svr.Serve(net))
}