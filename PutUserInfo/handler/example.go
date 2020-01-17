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
	example "gowork1_ihome/PutUserInfo/proto/example"
	"strconv"
	"time"
)

type Example struct {

}

func (e *Example) PutUserInfo(ctx context.Context,req *example.Request,rsp *example.Response)error  {
	//打印被调用的函数
	beego.Info("---------------- PUT  /api/v1.0/user/name PutUersinfo() ------------------")
	//创建返回空间
	rsp.Errno= utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	/*得到用户发送过来的name*/
	redis_confs := map[string]string{
		"key":utils.G_server_name,
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
		"password": utils.G_redis_passwd,
	}
	redis_conf,_ :=json.Marshal(redis_confs)
	bm, err := cache.NewCache("redis", string(redis_conf))
	if err != nil {
		beego.Info("缓存创建失败",err)
		rsp.Errno  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return  nil
	}
	//拼接key
	sessioniduserid := req.Sessionid+ "user_id"
	//获取userid
	value_id := bm.Get(sessioniduserid)
	ids := string(value_id.([]byte))
	id,err := strconv.Atoi(ids)

	user:=models.User{Id:id,Name:req.Username}
	/*更新对应user_id的name字段的内容*/
	//创建数据库句柄
	o:= orm.NewOrm()
	//更新
	_ , err =o.Update(&user ,"name")
	if err !=nil{
		rsp.Errno= utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)

		return nil
	}
	/*更新session user_id*/
	sessionidname :=  req.Sessionid + "name"
	bm.Put(sessioniduserid,string(user.Id),time.Second*600)

	bm.Put(sessionidname,string(user.Name),time.Second*600)

	rsp.Username = user.Name
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


