package handler

import (
	"context"
	"github.com/astaxie/beego"
	example"gowork1_ihome/GetArea/proto/example"

	"gowork1_ihome/IhomeWeb/utils"


)

type Example struct {

}


func (e *Example) GetArea(ctx context.Context,req *example.Request,rsp *example.Response) error {

	beego.Info("请求地区成功")
	//初始化 错误码
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	/*1从缓存中获取数据*/
	//准备连接redis信息
	//{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
	return nil
}

func (e *Example) Stream(ctx context.Context,req *example.StreamingRequest, stream example.Example_StreamStream) error {

	return nil
}

func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {

	return nil

}





