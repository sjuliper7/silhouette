package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sjuliper7/silhouette/common/config"
	"github.com/sjuliper7/silhouette/common/protocs"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello World i'm Client")

	userService := connectToUserService()

	result, err := userService.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Println("[userService][List] Failed when remote function %+v", err)
	}
	fmt.Println(result)
}

func connectToUserService() protocs.UsersClient {
	userPort := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(userPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect to", userPort, err)
	}

	return protocs.NewUsersClient(conn)
}
