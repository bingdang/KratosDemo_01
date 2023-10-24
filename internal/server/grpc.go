package server

import (
	v1 "KratosDemo_01/api/helloworld/v1"
	"KratosDemo_01/api/verifyCode"
	"KratosDemo_01/internal/conf"
	"KratosDemo_01/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	greeter *service.GreeterService,
	verifyCoderService *service.VerifyCodeService, //append 依赖wire自动注入进来
	logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	verifyCode.RegisterVerifyCodeServer(srv, verifyCoderService) //参数为注入进来的struct
	return srv
}
