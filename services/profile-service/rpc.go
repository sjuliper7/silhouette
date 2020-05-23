package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	"github.com/sjuliper7/silhouette/commons/models"
	GrpcDelivery "github.com/sjuliper7/silhouette/services/profile-service/delivery/http/grpc"
	"github.com/sjuliper7/silhouette/services/profile-service/usecase"
	"google.golang.org/grpc"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func initRemoteProcedurCall(db *sqlx.DB, kafkaConsumer *kafka.Consumer, profileUc usecase.ProfileUseCase) *grpc.Server {
	profileServer := GrpcDelivery.NewProfileServer(profileUc)

	svr := grpc.NewServer()
	models.RegisterProfilesServer(svr, profileServer)
	logrus.Infof("starting RPC server at %v", config.SERVICE_PROFILE_PORT)

	return svr
}
