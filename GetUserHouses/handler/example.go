package handler
import (
	"context"
	"github.com/micro/go-log"
	example "gowork1_ihome/GetUserHouses/proto/example"
	"github.com/astaxie/beego"
	"gowork1_ihome/IhomeWeb/utils"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"reflect"
	"gowork1_ihome/IhomeWeb/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)
type Example struct{}

func (e *Example) GetUserHouses(ctx context.Context, req *example.Request, rsp *example.Response) error {
	//打印被调用的函数
	beego.Info("获取当前用户所发布的房源 GetUserHouses /api/v1.0/user/houses")
	//创建返回空间
	rsp.Errno  =  utils.RECODE_OK
	rsp.Errmsg  = utils.RecodeText(rsp.Errno)
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
	sessioniduserid := req.Sessionid+ "user_id"
	//获取userid
	value_id := bm.Get(sessioniduserid)
	ids := string(value_id.([]byte))
	id,err := strconv.Atoi(ids)


	house_list :=[]models.House{}
	//创建数据库句柄
	o:=orm.NewOrm()
	qs:=o.QueryTable("house")
	num ,err :=qs.Filter("user_id",id).All(&house_list)
	if err !=nil{
		rsp.Errno  =  utils.RECODE_DBERR
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
	}
	if num==0 {
		rsp.Errno  =  utils.RECODE_NODATA
		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
	}
	house ,err :=json.Marshal(house_list)
	rsp.Mix= house
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








