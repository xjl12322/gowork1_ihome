package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
	"gowork1_ihome/IhomeWeb/handler"
	_ "gowork1_ihome/IhomeWeb/models"
	_ "gowork1_ihome/IhomeWeb/utils"
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
	//获取首页轮播图
	rou.GET("/api/v1.0/house/index",handler.GetIndex)
	//获取验证码图片
	rou.GET("/api/v1.0/imagecode/:uuid",handler.GetImageCd)
	//获取短信验证码
	rou.GET("/api/v1.0/smscode/:mobile",handler.GetSmscd)
	//获取短信验证码
	rou.POST("/api/v1.0/users",handler.PostRet)
	//获取session
	rou.GET("/api/v1.0/session",handler.GetSession)
	//登陆
	rou.POST("/api/v1.0/sessions",handler.PostLogin)
	//退出登陆
	rou.DELETE("/api/v1.0/session",handler.DeleteSession)
	//获取用户信息
	rou.GET("/api/v1.0/user",handler.GetUserInfo)
	//上传头像 POST
	rou.POST("/api/v1.0/user/avatar",handler.PostAvatar)
	//请求更新用户名 PUT
	rou.PUT("/api/v1.0/user/name",handler.PutUserInfo)
	//实名认证检查 GET
	rou.GET("/api/v1.0/user/auth",handler.GetUserAuth)
	////实名认证 post
	rou.POST("/api/v1.0/user/auth",handler.PostUserAuth)
	//请求当前用户已发布房源信息  GET
	rou.GET("/api/v1.0/user/houses",handler.GetUserHouses)


	service.Handle("/", rou)
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}