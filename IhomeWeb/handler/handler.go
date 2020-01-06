package handler

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro"
	GETAREA "gowork1_ihome/GetArea/proto/example"
	"net/http"
)

//获取地区信息
func GetArea(w http.ResponseWriter, r *http.Request,params httprouter.Params)  {


	beego.Info("请求地区信息 GetArea api/v1.0/areas")
	//创建服务获取句柄
	server := micro.NewService()
	//服务初始化
	server.Init()
	//调用服务返回句柄
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea",server.Client())
	//调用服务返回数据
	rsp,err := exampleClient.GetArea(context.TODO(),&GETAREA.Request{})
	if err != nil {
		http.Error(w,err.Error(),500)
		return
	}
	//接收数据
	//准备接收切片
	fmt.Println(rsp)

}






