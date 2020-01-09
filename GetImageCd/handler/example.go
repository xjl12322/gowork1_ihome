package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"github.com/micro/go-log"
	example "gowork1_ihome/GetImageCd/proto/example"
	"gowork1_ihome/IhomeWeb/utils"
	"image/color"
	"time"
)
type Example struct {

}

func (e *Example) GetImageCd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("获取验证码图片 GetImageCd /api/v1.0/imagecode/:uuid")
	/*生成验证码图片*/
	//创建图片句柄
	cap := captcha.New()
	//设置字体
	if err := cap.SetFont(utils.G_ziti); err != nil {
		panic(err.Error())
	}
	//设置图片大小
	cap.SetSize(91, 41)
	//设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)
	//设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	//生存随即的验证码图片
	img, str := cap.Create(4, captcha.NUM)

	/*将uuid和随即验证码进行缓存*/
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
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
	}
	//验证码与uuid进行缓存
	bm.Put(req.Uuid,str,time.Second*300)
	//图片解引用
	img1 := *img
	img2 := *img1.RGBA
	//返回错误信息
	rsp.Error= utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	//返回图片拆分
	rsp.Pix = []byte(img2.Pix)
	rsp.Stride = int64(img2.Stride)
	rsp.Max = &example.Response_Point{X:int64(img2.Rect.Max.X),Y:int64(img2.Rect.Max.Y)}
	rsp.Min = &example.Response_Point{X:int64(img2.Rect.Min.X),Y:int64(img2.Rect.Min.Y)}
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

