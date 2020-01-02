package config

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/commons/config"
	"github.com/sjuliper7/silhouette/commons/models"
	grpc_ "github.com/sjuliper7/silhouette/services/user-service/delivery/grpc"
	"github.com/sjuliper7/silhouette/services/user-service/delivery/rest"
	"github.com/sjuliper7/silhouette/services/user-service/repositories/mysql"
	"github.com/sjuliper7/silhouette/services/user-service/repositories/services"
	"github.com/sjuliper7/silhouette/services/user-service/usecase"
	"google.golang.org/grpc"
	"log"
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
		log.Println("Error when to connect grpc to profile service")
	}

	usecase := usecase.NewUserUsecase(userRepo, profileRepo)

	svr := grpc.NewServer()
	userServer := grpc_.NewUserServer(usecase)

	models.RegisterUsersServer(svr, userServer)
	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	//next running the to http
	net, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatalln(svr.Serve(net))
}

func initRestService(cg *Config) {
	userRepo := mysql.NewUserMysqlRepository(cg.DB)
	profileRepo, err := services.NewProfileRepository()
	if err != nil {
		logrus.Println("Error when to connect grpc to profile service")
	}

	usecase := usecase.NewUserUsecase(userRepo, profileRepo)

	router := mux.NewRouter()
	userRest := rest.NewUserServerRest(usecase)
	router.HandleFunc("/users", userRest.Resource).Methods("GET", "POST")
	router.HandleFunc("/users/{id}", userRest.Resource).Methods("GET", "PUT", "DELETE")

	logrus.Println("Starting Rest API at", config.REST_USER_PORT)

	http.ListenAndServe(config.REST_USER_PORT, router)

}
