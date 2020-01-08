package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-micro"
	GETAREA "gowork1_ihome/GetArea/proto/example"
	GETIMAGECD"gowork1_ihome/GetImageCd/proto/example"
	"gowork1_ihome/IhomeWeb/models"
	"image"
	"net/http"
	"gowork1_ihome/IhomeWeb/utils"

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
	area_list := []models.Area{}
	//循环接收数据
	for _,value := range rsp.Data{
		tmp := models.Area{Id:int(value.Aid),Name:value.Aname}
		area_list = append(area_list,tmp)
	}

	response := map[string]interface{}{
		"errno":rsp.Error,
		"errmsg:":rsp.Errmsg,
		"data":area_list,
	}
	//会传数据的时候三直接发送过去的并没有设置数据格式
	w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(response);err != nil{
		http.Error(w, err.Error(), 500)
		return
	}

}

//获取session信息
func GetSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	beego.Info("获取session信息 GetSession /api/v1.0/session")
	// we want to augment the response
	response := map[string]interface{}{
		"errno":  utils.RECODE_SESSIONERR,
		"errmsg": utils.RecodeText(utils.RECODE_SESSIONERR),
	}
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

//获取首页轮播图
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	beego.Info("获取首页轮播图 GetIndex api/v1.0/houses/index")
	response := map[string]interface{}{
		"errno": utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),
	}
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


}


func GetImageCd(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	beego.Info("获取验证码图片 GetImageCd /api/v1.0/imagecode/:uuid")

	//创建服务
	server:=micro.NewService()
	server.Init()
	// 调用服务
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd",server.Client())

	//获取uuid
	uuid:= ps.ByName("uuid")

	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid:uuid,
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//接收图片信息的 图片格式
	var img image.RGBA

	img.Stride = int(rsp.Stride)
	img.Pix =[]uint8(rsp.Pix)
	img.Rect.Min.X =int(rsp.Min.X)
	img.Rect.Min.Y =int(rsp.Min.Y)
	img.Rect.Max.X =int(rsp.Max.X)
	img.Rect.Max.Y =int(rsp.Max.Y)





}



