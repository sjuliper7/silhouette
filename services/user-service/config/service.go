package config

import (
	"github.com/gorilla/mux"
	"github.com/sjuliper7/silhouette/common/config"
	"github.com/sjuliper7/silhouette/common/protocs"
	"github.com/sjuliper7/silhouette/services/user-service/delivery"
	"github.com/sjuliper7/silhouette/services/user-service/repositories"
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
	repo := repositories.NewMysqlRepository(cg.DB)
	usecase := usecase.NewUserUsecase(repo)

	svr := grpc.NewServer()
	userServer := delivery.NewUserServer(usecase)

	protocs.RegisterUsersServer(svr, userServer)
	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	//next running the to http
	net, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatalln(svr.Serve(net))
}

func initRestService(cg *Config) {
	repo := repositories.NewMysqlRepository(cg.DB)
	usecase := usecase.NewUserUsecase(repo)

	router := mux.NewRouter()
	userRest := delivery.NewUserServerRest(usecase)
	router.HandleFunc("/users", userRest.Resource).Methods("GET", "POST")
	router.HandleFunc("/users/{id}", userRest.Resource).Methods("GET", "PUT", "DELETE")

	log.Println("Starting Rest API at", config.REST_USER_PORT)

	http.ListenAndServe(config.REST_USER_PORT, router)

}
