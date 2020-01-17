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
	example "gowork1_ihome/PostUserAuth/proto/example"
	"strconv"
	"time"
)
type Example struct{}



//实名认证
func (e *Example) PostUserAuth(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//打印被调用的函数
	beego.Info(" 实名认证 Postuserauth  api/v1.0/user/auth ")
	//创建返回空间
	rsp.Errno= utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	/*从session中获取我们的user_id*/
	//构建连接缓存的数据
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
	//拼接key
	sessioniduserid :=  req.Sessionid + "user_id"
	value_id := bm.Get(sessioniduserid)
	ids := string(value_id.([]byte))
	id,err := strconv.Atoi(ids)

	//创建user对象
	user := models.User{Id:id,
		Real_name:req.RealName,
		Id_card:req.IdCard,
	}
	o :=orm.NewOrm()
	//更新表
	_,err  =o.Update(&user ,"real_name","id_card")
	if err !=nil {
		rsp.Errno  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
		return nil
	}
	/*更新我们的session中的user_id*/
	bm.Put(sessioniduserid,string(user.Id),time.Second*600)
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








