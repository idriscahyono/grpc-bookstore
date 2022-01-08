package main

import (
	"github.com/idriscahyono/grpc-bookstore/app/configs"
	"github.com/idriscahyono/grpc-bookstore/app/grpc-server"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	configs.InitDB()
	grpc_server.InitGrpcServer()
}
