package main

import (
	"github.com/micro/go-micro"
	"gowork1_ihome/GetArea/handler"
	example"gowork1_ihome/GetArea/proto/example"


)

func main()  {
	//创建1个新的web服务
	service := micro.NewService(
		micro.Name("go.micro.srv.GetArea"),
		micro.Version("latest"),
	)
	service.Init()

	example.RegisterExampleHandler(service.Server(),new(handler.Example))






}

