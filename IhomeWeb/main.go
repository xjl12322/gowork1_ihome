package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"net/http"
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

	// register call handler

	//使用路由中间件来映射页面
	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("/gowork1_ihome/IhomeWeb/html"))

	// register html handler
	//映射前端页面
	service.Handle("/", rou)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}