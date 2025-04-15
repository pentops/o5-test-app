package service

import (

	// Replace with the actual path to the generated protobuf package

	"context"
	"database/sql"
	"time"

	sq "github.com/elgris/sqrl"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_pb"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_spb"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TestAccess struct {
	db sqrlx.Transactor
	test_spb.UnsafeEventLogServiceServer
}

func NewTestAccess(db sqrlx.Transactor) (*TestAccess, error) {
	return &TestAccess{
		db: db,
	}, nil
}

func (ta *TestAccess) RegisterGRPC(server grpc.ServiceRegistrar) {
	test_spb.RegisterEventLogServiceServer(server, ta)
}

func (ta *TestAccess) GetMessages(ctx context.Context, req *test_spb.GetMessagesRequest) (*test_spb.GetMessagesResponse, error) {

	_, err := j5auth.GetAuthenticatedAction(ctx)
	if err != nil {
		return nil, err
	}

	qq := sq.Select("message_id", "greeting_id", "event_id", "timestamp").
		From("greeting_message")

	if req.EventId == nil && req.GreetingId == nil {
		return nil, status.Error(codes.InvalidArgument, "Either EventId or GreetingId must be provided")
	}

	if req.EventId != nil {
		qq.Where("event_id = ?", req.EventId)
	}
	if req.GreetingId != nil {
		qq.Where("greeting_id = ?", req.GreetingId)
	}

	res := &test_spb.GetMessagesResponse{}

	err = ta.db.Transact(ctx, &sqrlx.TxOptions{
		ReadOnly:  true,
		Retryable: true,
		Isolation: sql.LevelDefault,
	}, func(ctx context.Context, tx sqrlx.Transaction) error {
		rows, err := tx.Select(ctx, qq)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			message := &test_pb.Message{}
			var timestamp time.Time
			if err := rows.Scan(&message.MessageId, &message.GreetingId, &message.EventId, &timestamp); err != nil {
				return err
			}
			message.Timestamp = timestamppb.New(timestamp)
			res.Messages = append(res.Messages, message)
		}
		if err := rows.Err(); err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
