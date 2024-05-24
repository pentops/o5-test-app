package service

import (
	"context"
	"database/sql"

	sq "github.com/elgris/sqrl"
	"github.com/pentops/o5-test-app/gen/test/v1/test_spb"
	"github.com/pentops/o5-test-app/gen/test/v1/test_tpb"
	"github.com/pentops/outbox.pg.go/outbox"
	"github.com/pentops/sqrlx.go/sqrlx"
)

type TestService struct {
	db *sqrlx.Wrapper
	*test_spb.UnimplementedTestServiceServer
}

func NewTestService(conn sqrlx.Connection) (*TestService, error) {
	db, err := sqrlx.New(conn, sq.Dollar)
	if err != nil {
		return nil, err
	}

	return &TestService{
		db: db,
	}, nil
}

func (ss *TestService) Hello(ctx context.Context, req *test_spb.HelloRequest) (*test_spb.HelloResponse, error) {

	msg := &test_tpb.GreetingMessage{
		Name: req.Name,
	}
	err := ss.db.Transact(ctx, &sqrlx.TxOptions{
		ReadOnly:  false,
		Retryable: true,
		Isolation: sql.LevelReadCommitted,
	}, func(ctx context.Context, tx sqrlx.Transaction) error {
		return outbox.Send(ctx, tx, msg)
	})
	if err != nil {
		return nil, err
	}

	return &test_spb.HelloResponse{
		Message:    "Hello, " + req.Name,
		AppVersion: "test-app-version",
	}, nil
}
