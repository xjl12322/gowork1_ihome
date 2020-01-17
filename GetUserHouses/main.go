package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gowork1_ihome/GetUserHouses/handler"
	example "gowork1_ihome/GetUserHouses/proto/example"

)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetUserHouses"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetUserHouses", service.Server(), new(subscriber.Example))
	//
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetUserHouses", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
