package main

import (
	"context"
	"database/sql"
	"fmt"
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

type DBConfig struct {
	PostgresURL  string `env:"POSTGRES_URL"`
	MaxOpenConns int    `env:"PG_MAX_CONNS" default:"5"`
	PingTimeout  int    `env:"PG_PING_TIMEOUT_SECONDS" default:"10"`
}

func (c *DBConfig) OpenDatabase(ctx context.Context) (*sql.DB, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(c.PingTimeout))

	defer cancel()

	db, err := sql.Open("postgres", c.PostgresURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(c.MaxOpenConns)

	for {
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("waiting for database: %w", err)
		}
		if err := db.PingContext(ctx); err != nil {
			log.WithError(ctx, err).Error("pinging PG")
			time.Sleep(time.Second)
			continue
		}
		break
	}

	return db, nil
}

func runMigrate(ctx context.Context, config struct {
	DBConfig
	MigrationsDir string `env:"MIGRATIONS_DIR" default:"./ext/db"`
}) error {
	db, err := config.OpenDatabase(ctx)
	if err != nil {
		return err
	}

	return goose.Up(db, config.MigrationsDir)
}

func runServe(ctx context.Context, config struct {
	ServeAddr string `env:"SERVE_ADDR" default:":8080"`
	DBConfig
}) error {

	db, err := config.OpenDatabase(ctx)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		service.GRPCMiddleware(version)...,
	)))

	if err := service.RegisterGRPC(grpcServer, db, version); err != nil {
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
