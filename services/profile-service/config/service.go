package config

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	"github.com/sjuliper7/silhouette/commons/models"
	grpc2 "github.com/sjuliper7/silhouette/services/profile-service/delivery/grpc"
	"github.com/sjuliper7/silhouette/services/profile-service/repositories/mysql"
	"google.golang.org/grpc"
	"net"
	kafkaC "github.com/sjuliper7/silhouette/services/profile-service/delivery/kafka"

	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
)

func (cf *Config) initService() {
	repo := mysql.NewMysqlProfileRepository(cf.DB)
	profileUc := usecase.NewProfileUseCase(repo)

	svr := grpc.NewServer()
	profileServer := grpc2.NewProfileServer(profileUc)
	//
	models.RegisterProfilesServer(svr, profileServer)
	logrus.Infof("starting RPC server at %v", config.SERVICE_PROFILE_PORT)

	err := kafkaC.ConsumeHandler(cf.KafkaConsumer, profileUc)
	if err != nil{
		logrus.Println("[config][service] failed to start consuming from kafka")
	}

	//next running the to http
	net, err := net.Listen("tcp", config.SERVICE_PROFILE_PORT)
	if err != nil {
		logrus.Fatalln("could not listen to %s: %v", config.SERVICE_PROFILE_PORT, err)
	}

	logrus.Fatalln(svr.Serve(net))
}
