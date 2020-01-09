package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"gowork1_ihome/GetSmscd/handler"
	example "gowork1_ihome/GetSmscd/proto/example"
)

func main()  {

	service := micro.NewService(
		micro.Name("go.micro.srv.GetSmscd"),
		micro.Version("latest"),
	)
	service.Init()
	example.RegisterExampleHandler(service.Server(),new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
