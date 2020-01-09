package main

import (
	"github.com/micro/go-micro"

	"gowork1_ihome/GetArea/handler"
	example "gowork1_ihome/GetArea/proto/example"
	"log"
)

func main()  {
	//创建1个新的web服务
	service := micro.NewService(
		micro.Name("go.micro.srv.GetArea"),
		micro.Version("latest"),
	)
	service.Init()

	example.RegisterExampleHandler(service.Server(),new(handler.Example))

	//micro.RegisterSubscriber("go.micro.srv.GetArea",service.Server(),new(subscriber.Example))


	if err := service.Run();err != nil{
		log.Fatal(err)
	}


}

