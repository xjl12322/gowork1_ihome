package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"github.com/micro/go-log"
	example "gowork1_ihome/GetSmscd/proto/example"
	"gowork1_ihome/IhomeWeb/models"
	"gowork1_ihome/IhomeWeb/utils"
	"math/rand"
	"reflect"
	"time"
)

type Example struct {

}

func (e *Example) GetSmscd(ctx context.Context,req *example.Request, rsp *example.Response) error {
	//初始化返回值

	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	/*验证手机号是否存在*/
	//创建数据库orm句柄
	o:=orm.NewOrm()
	//使用手机号作为查询条件
	user := models.User{Mobile:req.Mobile}

	err := o.Read(&user)
	//如果不报错就说明查找到了
	//查找到就说明手机号存在
	if err==nil{
		beego.Info("用户以存在")
		rsp.Error = utils.RECODE_MOBILEERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	/*验证图片验证码是否正确*/
	//连接redis
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
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//通过uuid查找图片验证码的值进行对比
	value:=bm.Get(req.Uuid)
	if value ==nil{
		beego.Info("redis获取失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//reflect.TypeOf(value)会返回当前数据的变量类型
	beego.Info(reflect.TypeOf(value),value)
	//格式转换
	value_str,_:= redis.String(value,nil)

	if value_str != req.Imagestr{
		beego.Info("数据不匹配 图片验证码值错误")
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	/*调用 短信接口发送短信*/
	//创建随机数
	r :=rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(9999)+1001
	beego.Info("验证码",size)
	err=bm.Put(req.Mobile,size,time.Second*300)
	if err!=nil {
		beego.Info("redis创建失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

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



