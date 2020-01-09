module gowork1_ihome

go 1.12

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/astaxie/beego v1.12.0
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/julienschmidt/httprouter v1.3.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v0.23.0
	github.com/micro/go-web v1.0.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	golang.org/x/image v0.0.0-20191214001246-9130b4cfad52 // indirect
	google.golang.org/appengine v1.6.0 // indirect
	google.golang.org/grpc v1.26.0 // indirect

)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191219195013-becbf705a915
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181219222714-6e267b5cc78e
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.0.0-20181220000619-583d854617af
	google.golang.org/appengine => github.com/golang/appengine v1.3.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20181219182458-5a97ab628bfb

)
