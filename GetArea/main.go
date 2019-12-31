package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-web"
)

func main()  {
	//创建1个新的web服务
	service := micro.NewService(
		micro.Name("go.micro.srv.GetArea"),
		micro.Version("latest"),
	)

}

