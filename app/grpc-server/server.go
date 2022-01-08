package grpc_server

import (
	"fmt"
	"github.com/idriscahyono/grpc-bookstore/app/proto/book"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitGrpcServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.Get("GRPC_PORT").(int)))

	if err != nil {
		log.Fatalln(err)
	}

	bookServer := book.Server{}

	grpcServer := grpc.NewServer()

	book.RegisterBookServiceServer(grpcServer, &bookServer)
	log.Printf("Server Running At %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
