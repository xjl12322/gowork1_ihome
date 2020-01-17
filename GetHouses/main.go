package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gowork1_ihome/GetHouses/handler"
	example "gowork1_ihome/GetHouses/proto/example"

)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetHouses"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))
	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetHouses", service.Server(), new(subscriber.Example))

	//// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.GetHouses", service.Server(), subscriber.Handler)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
