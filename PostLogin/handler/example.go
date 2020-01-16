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
	"gowork1_ihome/IhomeWeb/models"
	"gowork1_ihome/IhomeWeb/utils"
	example "gowork1_ihome/PostLogin/proto/example"
	"time"
)

type Example struct{}

func (e *Example) PostLogin(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("登陆  PostLogin /api/v1.0/sessions")
	rsp.Errno= utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	/*查询数据*/
	//创建数据库orm句柄
	o:=orm.NewOrm()
	//创建user对象
	var user models.User
	//创建查询条件句柄
	qs:=o.QueryTable("user")
	//通过qs句柄进行查询
	err :=qs.Filter("mobile",req.Mobile).One(&user)
	if err !=nil{
		beego.Info("查询数据失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	/*密码的校验*/
	if utils.Md5String(req.Password) != user.Password_hash{
		beego.Info("密码错误")
		rsp.Errno = utils.RECODE_PWDERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	/*创建sessionid 顺便就把数据返回*/
	sessionid :=utils.Md5String(req.Mobile+req.Password)

	rsp.Sessionid=sessionid
	/*将登陆信息进行缓存*/
	//配置缓存参数
	redis_conf := map[string]string{
		"key":utils.G_server_name,
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
		"password": utils.G_redis_passwd,
	}
	beego.Info(redis_conf)
	//将map进行转化成为json
	redis_conf_js,_ :=json.Marshal(redis_conf)
	//创建redis句柄
	bm ,err :=cache.NewCache("redis",string(redis_conf_js))
	if err!=nil{
		beego.Info("redis连接失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}
	/*拼接key*/
	//user_id
	sessionuser_id := sessionid+"user_id"
	bm.Put(sessionuser_id,user.Id,time.Second*1000)
	//name
	sessionname := sessionid+"name"
	bm.Put(sessionname,user.Name,time.Second*1000)
	//mobile
	sessionmobile := sessionid+"mobile"
	bm.Put(sessionmobile,user.Mobile,time.Second*1000)
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




