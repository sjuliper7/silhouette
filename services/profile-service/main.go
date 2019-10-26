package main

import "github.com/sjuliper7/silhouette/services/profile-service/config"

func init() {
	config.LoadConfig()
}

func main() {
	config.InitConfig()
}
