// Code generated by protoc-gen-go-o5-messaging. DO NOT EDIT.
// versions:
// - protoc-gen-go-o5-messaging 0.0.0
// source: test/v1/topic/test_topic.proto

package test_tpb

import (
	context "context"

	o5msg "github.com/pentops/o5-messaging/o5msg"
)

// Service: TestTopic
type TestTopicTxSender[C any] struct {
	sender o5msg.TxSender[C]
}

func NewTestTopicTxSender[C any](sender o5msg.TxSender[C]) *TestTopicTxSender[C] {
	sender.Register(o5msg.TopicDescriptor{
		Service: "test.v1.topic.TestTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "Greeting",
				Message: (*GreetingMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &TestTopicTxSender[C]{sender: sender}
}

type TestTopicCollector[C any] struct {
	collector o5msg.Collector[C]
}

func NewTestTopicCollector[C any](collector o5msg.Collector[C]) *TestTopicCollector[C] {
	collector.Register(o5msg.TopicDescriptor{
		Service: "test.v1.topic.TestTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "Greeting",
				Message: (*GreetingMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &TestTopicCollector[C]{collector: collector}
}

type TestTopicPublisher struct {
	publisher o5msg.Publisher
}

func NewTestTopicPublisher(publisher o5msg.Publisher) *TestTopicPublisher {
	publisher.Register(o5msg.TopicDescriptor{
		Service: "test.v1.topic.TestTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "Greeting",
				Message: (*GreetingMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &TestTopicPublisher{publisher: publisher}
}

// Method: Greeting

func (msg *GreetingMessage) O5MessageHeader() o5msg.Header {
	header := o5msg.Header{
		GrpcService:      "test.v1.topic.TestTopic",
		GrpcMethod:       "Greeting",
		Headers:          map[string]string{},
		DestinationTopic: "o5-test-topic",
	}
	return header
}

func (send TestTopicTxSender[C]) Greeting(ctx context.Context, sendContext C, msg *GreetingMessage) error {
	return send.sender.Send(ctx, sendContext, msg)
}

func (collect TestTopicCollector[C]) Greeting(sendContext C, msg *GreetingMessage) {
	collect.collector.Collect(sendContext, msg)
}

func (publish TestTopicPublisher) Greeting(ctx context.Context, msg *GreetingMessage) {
	publish.publisher.Publish(ctx, msg)
}
