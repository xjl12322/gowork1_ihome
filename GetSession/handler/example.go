package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"gowork1_ihome/IhomeWeb/utils"

	example "gowork1_ihome/GetSession/proto/example"
)
type Example struct{}

func (e *Example) GetSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取session信息 GetSession /api/v1.0/session")
	//初始化返回值
	rsp.Errno = utils.RECODE_OK;
	rsp.Errmsg =utils.RecodeText(rsp.Errno)

	/*获取usernamer*/
	//配置缓存参数
	redis_conf := map[string]string{
		"key":utils.G_server_name,
		//127.0.0.1:6379
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

	username := bm.Get(req.Sessionid+"name")
	/*没有返回失败*/
	if username == nil{
		beego.Info("获取数据并不存在",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}
	/*有就返回成功*/
	rsp.UserName,_ = redis.String(username,nil)
	return nil
}






















