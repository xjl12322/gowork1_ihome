package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gowork1_ihome/PostUserAuth/handler"
	example "gowork1_ihome/PostUserAuth/proto/example"
)
func main() {
	// New Service
	service :=micro.NewService(
		micro.Name("go.micro.srv.PostUserAuth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))


	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}