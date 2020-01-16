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
	example "gowork1_ihome/PostAvatar/proto/example"
	"io/ioutil"
	"path"
	"reflect"
	"strconv"
)


type Example struct {

}

func (e *Example) PostAvatar(ctx context.Context,req *example.Request,rsp *example.Response) error {
	beego.Info("上传用户头像 PostAvatar /api/v1.0/user/avatar")
	//初始化返回正确的返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//检查下数据是否正常
	beego.Info(len(req.Avatar),req.Filesize)
	/*获取文件的后缀名*/     //dsnlkjfajadskfksda.sadsdasd.sdasd.jpg
	fileext := path.Ext(req.Filename)
	//上传数据

	//fh,err := os.Open(req.Sessionid+fileext)
	//if err != nil {
	//	fmt.Println("error opening file")
	//	return err
	//}
	//iocopy
	redis_conf := map[string]string{
		"key":utils.G_server_name,
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
		"password": utils.G_redis_passwd,
	}

	redis_config,_ :=json.Marshal(redis_conf)
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
	//3792d84071ad061b945315a92d8520b8user_id
	//  3792d84071ad061b945315a92d8520b8user_id
	//[53]

	beego.Info(value_id,reflect.TypeOf(value_id))
	//数据格式转换
	ids := string(value_id.([]byte))
	id,err := strconv.Atoi(ids)
	//"http://127.0.0.1:10086/user/"
	ioutil.WriteFile("D:/golands/ihome/gowork1_ihome/IhomeWeb/html/user/"+req.Sessionid+fileext,req.Avatar, 0777)
	//_, err = io.Copy([]byte(req.Avatar), fh)
	//if err != nil {
	//	return err
	//}

	user := models.User{Id:id,Avatar_url:"http://"+utils.G_fastdfs_addr+":"+utils.G_fastdfs_port+"/user/"+req.Sessionid+fileext}
	/*将当前fastdfs-url 存储到我们当前用户的表中*/
	o:=orm.NewOrm()
	_ ,err =o.Update(&user ,"avatar_url")
	if err !=nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
	}
	rsp.AvatarUrl = user.Avatar_url
	return nil
}


//func (e *Example) PostAvatar(ctx context.Context, req *example.Request, rsp *example.Response) error {
//	beego.Info("上传用户头像 PostAvatar /api/v1.0/user/avatar")
//	//初始化返回正确的返回值
//	rsp.Errno = utils.RECODE_OK
//	rsp.Errmsg = utils.RecodeText(rsp.Errno)
//
//	//检查下数据是否正常
//	beego.Info(len(req.Avatar),req.Filesize)
//
//
//	/*获取文件的后缀名*/     //dsnlkjfajadskfksda.sadsdasd.sdasd.jpg
//	beego.Info("后缀名",path.Ext(req.Filename))
//
//	/*存储文件到fastdfs当中并且获取 url*/
//	//.jpg
//	fileext :=path.Ext(req.Filename)
//	//group1 group1/M00/00/00/wKgLg1t08pmANXH1AAaInSze-cQ589.jpg
//	//上传数据
//	Group,FileId ,err :=  models.UploadByBuffer(req.Avatar,fileext[1:])
//	if err != nil {
//		beego.Info("Postupavatar  models.UploadByBuffer err" ,err)
//		rsp.Errno = utils.RECODE_IOERR
//		rsp.Errmsg = utils.RecodeText(rsp.Errno)
//		return nil
//	}
//	beego.Info(Group)
//
//	/*通过session 获取我们当前现在用户的uesr_id*/
//	redis_config_map := map[string]string{
//		"key":utils.G_server_name,
//		//"conn":"127.0.0.1:6379",
//		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
//		"dbNum":utils.G_redis_dbnum,
//	}
//	beego.Info(redis_config_map)
//	redis_config ,_:=json.Marshal(redis_config_map)
//	beego.Info( string(redis_config) )
//	//连接redis数据库 创建句柄
//	bm, err := cache.NewCache("redis", string(redis_config) )
//	if err != nil {
//		beego.Info("缓存创建失败",err)
//		rsp.Errno  =  utils.RECODE_DBERR
//		rsp.Errmsg  = utils.RecodeText(rsp.Errno)
//		return  nil
//	}
//	//拼接key
//	sessioniduserid :=  req.Sessionid + "user_id"
//
//	//获得当前用户的userid
//	value_id :=bm.Get(sessioniduserid)
//	beego.Info(value_id,reflect.TypeOf(value_id))
//
//
//	id :=  int(value_id.([]uint8)[0])
//	beego.Info(id ,reflect.TypeOf(id))
//
//	//创建表对象
//	user := models.User{Id:id,Avatar_url:FileId}
//	/*将当前fastdfs-url 存储到我们当前用户的表中*/
//	o:=orm.NewOrm()
//	//将图片的地址存入表中
//	_ ,err =o.Update(&user ,"avatar_url")
//	if err !=nil {
//
//		rsp.Errno = utils.RECODE_DBERR
//		rsp.Errmsg = utils.RecodeText(rsp.Errno)
//
//	}
//
//	//回传图片地址
//	rsp.AvatarUrl=FileId
//
//	return nil
//}
//









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











