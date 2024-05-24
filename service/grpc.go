package service

import (
	"github.com/pentops/go-grpc-helpers/grpcerror"
	"github.com/pentops/go-grpc-helpers/protovalidatemw"
	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/gen/test/v1/test_tpb"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
	"gopkg.daemonl.com/log"
	"gopkg.daemonl.com/log/grpc_log"
)

func GRPCMiddleware(version string) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_log.UnaryServerInterceptor(log.DefaultContext, log.DefaultTrace, log.DefaultLogger),
		grpcerror.UnaryServerInterceptor(log.DefaultLogger),
		protovalidatemw.UnaryServerInterceptor(),
	}
}

func RegisterGRPC(conn sqrlx.Connection, server *grpc.Server) error {

	testService, err := NewTestService(conn)
	if err != nil {
		return err
	}
	test_spb.RegisterTestServiceServer(server, testService)

	testWorker, err := NewTestWorker()
	if err != nil {
		return err
	}
	test_tpb.RegisterTestTopicServer(server, testWorker)

	return nil
}
