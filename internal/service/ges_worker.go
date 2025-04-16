package service

import (
	"context"
	"database/sql"

	sq "github.com/elgris/sqrl"
	"github.com/pentops/o5-test-app/internal/gen/test/v1/test_tpb"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PublishWorker struct {
	db sqrlx.Transactor
	test_tpb.UnsafeGreetingPublishTopicServer
}

func NewPublishWorker(db sqrlx.Transactor) (*PublishWorker, error) {
	return &PublishWorker{
		db: db,
	}, nil
}

func (ww *PublishWorker) RegisterGRPC(server grpc.ServiceRegistrar) {
	test_tpb.RegisterGreetingPublishTopicServer(server, ww)
}

func (ww *PublishWorker) GreetingEvent(ctx context.Context, req *test_tpb.GreetingEventMessage) (*emptypb.Empty, error) {

	message, err := j5auth.GetMessageCause(ctx)
	if err != nil {
		return nil, err
	}
	err = ww.db.Transact(ctx, &sqrlx.TxOptions{
		ReadOnly:  false,
		Retryable: true,
		Isolation: sql.LevelDefault,
	}, func(ctx context.Context, tx sqrlx.Transaction) error {

		_, err := tx.Insert(ctx, sq.
			Insert("greeting_message").
			Columns("greeting_id", "event_id", "message_id").
			Values(req.Keys.GreetingId, req.Metadata.EventId, message.MessageId))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
