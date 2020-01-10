package handler


import (
	"context"

	example "gowork1_ihome/DeleteSession/proto/example"
	"github.com/astaxie/beego"
	"gowork1_ihome/IhomeWeb/utils"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_"github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/garyburd/redigo/redis"

)

type Example struct{}

func (e *Example) DeleteSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("DeleteSession  退出登陆 /api/v1.0/session")
	//返回值初始化
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg =utils.RecodeText(rsp.Errno)

	/*连接redis*/
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
		rsp.Errno= utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}











}







