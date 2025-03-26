package main

import (
	"context"

	"github.com/pentops/grpc.go/grpcbind"
	"github.com/pentops/o5-test-app/internal/service"
	"github.com/pentops/runner/commander"
	"github.com/pentops/sqrlx.go/pgenv"
	"github.com/pressly/goose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var version string

func main() {
	mainGroup := commander.NewCommandSet()
	mainGroup.Add("serve", commander.NewCommand(runServe))
	mainGroup.Add("migrate", commander.NewCommand(runMigrate))
	mainGroup.RunMain("o5-test-app", version)
}

func runMigrate(ctx context.Context, config struct {
	pgenv.DatabaseConfig
	MigrationsDir string `env:"MIGRATIONS_DIR" default:"./ext/db"`
}) error {
	db, err := config.OpenPostgres(ctx)
	if err != nil {
		return err
	}

	return goose.Up(db, config.MigrationsDir)
}

func runServe(ctx context.Context, cfg struct {
	GRPCBind string `env:"GRPC_BIND" default:":8080"`
	pgenv.DatabaseConfig
}) error {

	db, err := cfg.OpenPostgresTransactor(ctx)
	if err != nil {
		return err
	}

	app, err := service.NewApp(db, version)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		service.GRPCMiddleware(version)...,
	))

	app.RegisterGRPC(grpcServer)
	reflection.Register(grpcServer)

	return grpcbind.ListenAndServe(ctx, grpcServer, cfg.GRPCBind)
}
