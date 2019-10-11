package main

import (
	"fmt"

	"github.com/sjuliper7/silhouette/services/user-service/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	fmt.Println("Hello")
	config.InitConfig()
}
