package main

import (
	"context"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pentops/runner/commander"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.daemonl.com/log"
)

var Version string

func main() {

	mainGroup := commander.NewCommandSet()

	mainGroup.Add("serve", commander.NewCommand(runServe))

	mainGroup.RunMain("testapp", Version)
}

func runServe(ctx context.Context, config struct {
	ServeAddr string `env:"SERVE_ADDR" default:":8080"`
}) error {

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		service.GRPCUnaryMiddleware(Version)...,
	)))

	if err := service.RegisterGRPC(grpcServer); err != nil {
		return err
	}

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.ServeAddr)
	if err != nil {
		return err
	}

	log.WithField(ctx, "addr", lis.Addr().String()).Info("server listening")
	closeOnContextCancel(ctx, grpcServer)

	return grpcServer.Serve(lis)
}

func closeOnContextCancel(ctx context.Context, srv *grpc.Server) {
	go func() {
		<-ctx.Done()
		srv.GracefulStop()
	}()
}
