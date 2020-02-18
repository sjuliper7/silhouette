package config

import (
	"github.com/gorilla/mux"
	"github.com/koinworks/asgard-heimdal/libs/logger"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	"github.com/sjuliper7/silhouette/commons/models"
	grpc_ "github.com/sjuliper7/silhouette/services/user-service/delivery/grpc"
	"github.com/sjuliper7/silhouette/services/user-service/delivery/rest"
	kafkaProducer "github.com/sjuliper7/silhouette/services/user-service/repositories/kafka"
	"github.com/sjuliper7/silhouette/services/user-service/repositories/mysql"
	"github.com/sjuliper7/silhouette/services/user-service/repositories/services"
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func (cf *Config) initService() {
	//initRpcService(cf)
	initRestService(cf)
}

func initRpcService(cg *Config) {
	userRepo := mysql.NewUserMysqlRepository(cg.DB)
	profileRepo, err := services.NewProfileRepository()
	if err != nil {
		logger.Errf("Error when to connect grpc to profile service, %v", err)
	}

	kafkaRepository := kafkaProducer.NewKafkaRepository(cg.KafkaProducer)

	usecase := usecase.NewUserUsecase(userRepo, profileRepo, kafkaRepository)

	svr := grpc.NewServer()
	userServer := grpc_.NewUserServer(usecase)

	models.RegisterUsersServer(svr, userServer)
	logrus.Infof("Starting RPC server at %v", config.SERVICE_USER_PORT)

	//next running the to http
	net, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		logrus.Errorf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	logrus.Fatalln(svr.Serve(net))
}

func initRestService(cg *Config) {

	kafkaRepo := kafkaProducer.NewKafkaRepository(cg.KafkaProducer)
	userRepo := mysql.NewUserMysqlRepository(cg.DB)
	profileRepo, err := services.NewProfileRepository()

	if err != nil {
		logrus.Infof("Error when to connect grpc to profile service %v", err)
	}

	profileUsecase := usecase.NewUserUsecase(userRepo, profileRepo, kafkaRepo)

	router := mux.NewRouter()
	userRest := rest.NewUserServerRest(profileUsecase)
	router.HandleFunc("/users", userRest.Resource).Methods("GET", "POST")
	router.HandleFunc("/users/{id}", userRest.Resource).Methods("GET", "PUT", "DELETE")

	logrus.Infof("Starting Rest API at %v", config.REST_USER_PORT)

	http.ListenAndServe(config.REST_USER_PORT, router)

}
