module gowork1_ihome

go 1.12

require (
	github.com/go-log/log v0.1.0
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.5.0

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
