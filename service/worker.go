package service

import (
	"context"

	"github.com/pentops/o5-test-app/gen/test/v1/test_tpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TestWorker struct {
	*test_tpb.UnimplementedTestTopicServer
}

func NewTestWorker() (*TestWorker, error) {
	return &TestWorker{}, nil
}

func (ww *TestWorker) Greeting(ctx context.Context, req *test_tpb.GreetingMessage) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
