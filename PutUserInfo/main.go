package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gowork1_ihome/PutUserInfo/handler"
	example "gowork1_ihome/PutUserInfo/proto/example"
)

func main()  {

	service := micro.NewService(
		micro.Name("go.micro.srv.PutUserInfo"),
		micro.Version("latest"),
		)

	service.Init()

	example.RegisterExampleHandler(service.Server(),new(handler.Example))

	if err := service.Run(); err !=nil {
		log.Fatal(err)
	}

}
