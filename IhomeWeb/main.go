package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"net/http"
	"gowork1_ihome/IhomeWeb/handler"
)

type H struct {

}

func (e *H)te(w http.ResponseWriter,r http.Request)  {
	fmt.Println("1111111111111111111111111111111111111111111")
}
func main()  {

	//创建1个新的web服务
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		web.Address(":10086"),

		)
	// initialise service
	//服务初始化

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register call GetArea

	//使用路由中间件来映射页面
	// register html GetArea
	//映射前端页面
	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("D:/golands/ihome/gowork1_ihome/IhomeWeb/html"))
	//rou := http.NotFoundHandler()
	//rou = http.FileServer(http.Dir("D:/golands/ihome/gowork1_ihome/IhomeWeb/html"))

	//获取地区请求

	rou.GET("/api/v1.0/areas",handler.GetArea)
	service.Handle("/", rou)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}