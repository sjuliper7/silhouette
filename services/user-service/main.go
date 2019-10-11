package main

import (
	"github.com/sjuliper7/silhouette/services/user-service/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	config.InitConfig()
}
