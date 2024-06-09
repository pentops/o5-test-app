package main

import (
	"context"
	"database/sql"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-test-app/internal/service"
	"github.com/pentops/runner/commander"
	"github.com/pressly/goose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var version string

func main() {

	mainGroup := commander.NewCommandSet()

	mainGroup.Add("serve", commander.NewCommand(runServe))
	mainGroup.Add("migrate", commander.NewCommand(runMigrate))

	mainGroup.RunMain("testapp", version)
}

func openDatabase(ctx context.Context, dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	for {
		if err := db.Ping(); err != nil {
			log.WithError(ctx, err).Error("pinging PG")
			time.Sleep(time.Second)
			continue
		}
		break
	}

	return db, nil
}
func runMigrate(ctx context.Context, config struct {
	MigrationsDir string `env:"MIGRATIONS_DIR" default:"./ext/db"`
	PostgresURL   string `env:"POSTGRES_URL"`
}) error {
	db, err := openDatabase(ctx, config.PostgresURL)
	if err != nil {
		return err
	}

	return goose.Up(db, config.MigrationsDir)
}

func runServe(ctx context.Context, config struct {
	ServeAddr   string `env:"SERVE_ADDR" default:":8080"`
	PostgresURL string `env:"POSTGRES_URL"`
}) error {

	db, err := openDatabase(ctx, config.PostgresURL)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		service.GRPCMiddleware(Version)...,
	)))

	if err := service.RegisterGRPC(grpcServer, db, Version); err != nil {
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
