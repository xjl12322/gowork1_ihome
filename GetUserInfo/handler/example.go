package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"github.com/micro/go-log"
	example "gowork1_ihome/GetUserInfo/proto/example"
	"gowork1_ihome/IhomeWeb/models"
	"gowork1_ihome/IhomeWeb/utils"
	"reflect"
	"strconv"

)

type Example struct {

}

func (e *Example) GetUserInfo(ctx context.Context,req *example.Request, rsp *example.Response) error {
	beego.Info("---------------- GET  /api/v1.0/user Getuserinfo() ------------------")
	//打印sessionid
	beego.Info(req.Sessionid,reflect.TypeOf(req.Sessionid))
	//错误码
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)

	redis_conf := map[string]string{
		"key":utils.G_server_name,
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
		"password": utils.G_redis_passwd,
	}
	beego.Info(redis_conf)
	//将map进行转化成为json
	redis_config,_ :=json.Marshal(redis_conf)

	//连接redis数据库 创建句柄
	bm, err := cache.NewCache("redis", string(redis_config))
	if err != nil {
		beego.Info("缓存创建失败",err)
		rsp.Errno  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return  nil
	}
	//拼接用户信息缓存字段
	sessioniduserid :=  req.Sessionid + "user_id"
	//获取到当前登陆用户的user_id
	value_id :=bm.Get(sessioniduserid)
	//打印
	beego.Info(value_id,reflect.TypeOf(value_id))
	//数据格式转换
	id := int(value_id.([]uint8)[0])
	beego.Info(id ,reflect.TypeOf(id))
	//创建user表
	user := models.User{Id:id}
	//查询表
	o := orm.NewOrm()
	err =o.Read(&user)
	if err !=nil{
		rsp.Errno  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return  nil
	}
	//将查询到的数据依次赋值
	rsp.UserId= user.Id
	rsp.Name= user.Name
	rsp.Mobile = user.Mobile
	rsp.RealName = user.Real_name
	rsp.IdCard = user.Id_card
	rsp.AvatarUrl = user.Avatar_url
	return nil

}


// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}



