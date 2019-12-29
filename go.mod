module gowork1_ihome

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/uuid v1.1.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/micro/cli v0.2.0 // indirect
	github.com/micro/go-log v0.1.0
	github.com/micro/go-web v1.0.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	golang.org/x/crypto v0.0.0-20190530122614-20be4c3c3ed5 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.6.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect

)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191219195013-becbf705a915
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181219222714-6e267b5cc78e
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.0.0-20181220000619-583d854617af
	google.golang.org/appengine => github.com/golang/appengine v1.3.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20181219182458-5a97ab628bfb
	google.golang.org/grpc => github.com/grpc/grpc-go v1.17.0

)
