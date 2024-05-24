package service

import (
	"github.com/interxfi/go-auth/auth"
	"github.com/pentops/go-grpc-helpers/grpcerror"
	"github.com/pentops/go-grpc-helpers/protovalidatemw"
	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
	"google.golang.org/grpc"
	"gopkg.daemonl.com/log"
	"gopkg.daemonl.com/log/grpc_log"
)

func GRPCMiddleware(version string) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_log.UnaryServerInterceptor(log.DefaultContext, log.DefaultTrace, log.DefaultLogger),
		grpcerror.UnaryServerInterceptor(log.DefaultLogger),
		protovalidatemw.UnaryServerInterceptor(),
		auth.GRPCMiddleware,
	}
}

func RegisterGRPC(server *grpc.Server) error {

	testService, err := NewTestService()
	if err != nil {
		return err
	}
	test_spb.RegisterTestServiceServer(server, testService)

	return nil
}
