package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gowork1_ihome/PostRet/handler"
	example "gowork1_ihome/PostRet/proto/example"
)
func main()  {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.PostRet"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()
	example.RegisterExampleHandler(service.Server(),new(handler.Example))
	//micro.RegisterSubscriber("go.micro.srv.PostRet", service.Server(), new(subscriber.Example))
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.PostRet", service.Server(), subscriber.Handler)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}




