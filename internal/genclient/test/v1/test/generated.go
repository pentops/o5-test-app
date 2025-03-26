package test

// Code generated by jsonapi. DO NOT EDIT.
// Source: github.com/pentops/o5-test-app/internal/genclient/test/v1/test

import (
	context "context"
	json "encoding/json"
	errors "errors"
	list "github.com/pentops/o5-test-app/internal/genclient/j5/list/v1/list"
	state "github.com/pentops/o5-test-app/internal/genclient/j5/state/v1/state"
	url "net/url"
	strings "strings"
)

type Requester interface {
	Request(ctx context.Context, method string, path string, body interface{}, response interface{}) error
}

// GreetingQueryService
type GreetingQueryService struct {
	Requester
}

func NewGreetingQueryService(requester Requester) *GreetingQueryService {
	return &GreetingQueryService{
		Requester: requester,
	}
}

func (s GreetingQueryService) GreetingGet(ctx context.Context, req *GreetingGetRequest) (*GreetingGetResponse, error) {
	pathParts := make([]string, 6)
	pathParts[0] = ""
	pathParts[1] = "test"
	pathParts[2] = "v1"
	pathParts[3] = "greeting"
	pathParts[4] = "q"
	if req.GreetingId == "" {
		return nil, errors.New("required field \"GreetingId\" not set")
	}
	pathParts[5] = req.GreetingId
	path := strings.Join(pathParts, "/")
	if query, err := req.QueryParameters(); err != nil {
		return nil, err
	} else if len(query) > 0 {
		path += "?" + query.Encode()
	}
	resp := &GreetingGetResponse{}
	err := s.Request(ctx, "GET", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s GreetingQueryService) GreetingList(ctx context.Context, req *GreetingListRequest) (*GreetingListResponse, error) {
	pathParts := make([]string, 5)
	pathParts[0] = ""
	pathParts[1] = "test"
	pathParts[2] = "v1"
	pathParts[3] = "greeting"
	pathParts[4] = "q"
	path := strings.Join(pathParts, "/")
	if query, err := req.QueryParameters(); err != nil {
		return nil, err
	} else if len(query) > 0 {
		path += "?" + query.Encode()
	}
	resp := &GreetingListResponse{}
	err := s.Request(ctx, "GET", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s GreetingQueryService) GreetingEvents(ctx context.Context, req *GreetingEventsRequest) (*GreetingEventsResponse, error) {
	pathParts := make([]string, 7)
	pathParts[0] = ""
	pathParts[1] = "test"
	pathParts[2] = "v1"
	pathParts[3] = "greeting"
	pathParts[4] = "q"
	if req.GreetingId == "" {
		return nil, errors.New("required field \"GreetingId\" not set")
	}
	pathParts[5] = req.GreetingId
	pathParts[6] = "events"
	path := strings.Join(pathParts, "/")
	if query, err := req.QueryParameters(); err != nil {
		return nil, err
	} else if len(query) > 0 {
		path += "?" + query.Encode()
	}
	resp := &GreetingEventsResponse{}
	err := s.Request(ctx, "GET", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GreetingGetRequest
type GreetingGetRequest struct {
	GreetingId string `json:"-" path:"greetingId"`
}

func (s GreetingGetRequest) QueryParameters() (url.Values, error) {
	values := url.Values{}
	return values, nil
}

// GreetingGetResponse
type GreetingGetResponse struct {
	Greeting *GreetingState `json:"greeting"`
}

// GreetingListRequest
type GreetingListRequest struct {
	Page  *list.PageRequest  `json:"-" query:"page"`
	Query *list.QueryRequest `query:"query" json:"-"`
}

func (s *GreetingListRequest) SetPageToken(pageToken string) {
	if s.Page == nil {
		s.Page = &list.PageRequest{}
	}
	s.Page.Token = &pageToken
}

func (s GreetingListRequest) QueryParameters() (url.Values, error) {
	values := url.Values{}
	if s.Page != nil {
		bb, err := json.Marshal(s.Page)
		if err != nil {
			return nil, err
		}
		values.Set("page", string(bb))
	}
	if s.Query != nil {
		bb, err := json.Marshal(s.Query)
		if err != nil {
			return nil, err
		}
		values.Set("query", string(bb))
	}
	return values, nil
}

// GreetingListResponse
type GreetingListResponse struct {
	Greeting []*GreetingState   `json:"greeting,omitempty"`
	Page     *list.PageResponse `json:"page,omitempty"`
}

func (s GreetingListResponse) GetPageToken() *string {
	if s.Page == nil {
		return nil
	}
	return s.Page.NextToken
}

func (s GreetingListResponse) GetItems() []*GreetingState {
	return s.Greeting
}

// GreetingEventsRequest
type GreetingEventsRequest struct {
	GreetingId string             `json:"-" path:"greetingId"`
	Page       *list.PageRequest  `json:"-" query:"page"`
	Query      *list.QueryRequest `json:"-" query:"query"`
}

func (s *GreetingEventsRequest) SetPageToken(pageToken string) {
	if s.Page == nil {
		s.Page = &list.PageRequest{}
	}
	s.Page.Token = &pageToken
}

func (s GreetingEventsRequest) QueryParameters() (url.Values, error) {
	values := url.Values{}
	if s.Page != nil {
		bb, err := json.Marshal(s.Page)
		if err != nil {
			return nil, err
		}
		values.Set("page", string(bb))
	}
	if s.Query != nil {
		bb, err := json.Marshal(s.Query)
		if err != nil {
			return nil, err
		}
		values.Set("query", string(bb))
	}
	return values, nil
}

// GreetingEventsResponse
type GreetingEventsResponse struct {
	Events []*GreetingEvent   `json:"events,omitempty"`
	Page   *list.PageResponse `json:"page,omitempty"`
}

func (s GreetingEventsResponse) GetPageToken() *string {
	if s.Page == nil {
		return nil
	}
	return s.Page.NextToken
}

func (s GreetingEventsResponse) GetItems() []*GreetingEvent {
	return s.Events
}

// GreetingCommandService
type GreetingCommandService struct {
	Requester
}

func NewGreetingCommandService(requester Requester) *GreetingCommandService {
	return &GreetingCommandService{
		Requester: requester,
	}
}

func (s GreetingCommandService) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	pathParts := make([]string, 7)
	pathParts[0] = ""
	pathParts[1] = "test"
	pathParts[2] = "v1"
	pathParts[3] = "greeting"
	pathParts[4] = "test"
	pathParts[5] = "v1"
	pathParts[6] = "echo"
	path := strings.Join(pathParts, "/")
	resp := &HelloResponse{}
	err := s.Request(ctx, "POST", path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HelloRequest
type HelloRequest struct {
	GreetingId  string     `json:"greetingId"`
	Name        string     `json:"name"`
	ThrowError  *TestError `json:"throwError,omitempty"`
	WorkerError *TestError `json:"workerError,omitempty"`
}

// HelloResponse
type HelloResponse struct {
	Greeting *GreetingState `json:"greeting"`
}

// GreetingEvent Proto: GreetingEvent
type GreetingEvent struct {
	Metadata   *state.EventMetadata `json:"metadata"`
	GreetingId string               `json:"greetingId"`
	Event      *GreetingEventType   `json:"event"`
}

// TestError Proto: TestError
type TestError struct {
	Message string `json:"message,omitempty"`
	Code    uint32 `json:"code,omitempty"`
}

// GreetingData Proto: GreetingData
type GreetingData struct {
	Name         string  `json:"name,omitempty"`
	ReplyMessage *string `json:"replyMessage,omitempty"`
	AppVersion   string  `json:"appVersion,omitempty"`
}

// GreetingEventType Proto Oneof: test.v1.GreetingEventType
type GreetingEventType struct {
	J5TypeKey string                       `json:"!type,omitempty"`
	Initiated *GreetingEventType_Initiated `json:"initiated,omitempty"`
	Replied   *GreetingEventType_Replied   `json:"replied,omitempty"`
}

func (s GreetingEventType) OneofKey() string {
	if s.Initiated != nil {
		return "initiated"
	}
	if s.Replied != nil {
		return "replied"
	}
	return ""
}

func (s GreetingEventType) Type() interface{} {
	if s.Initiated != nil {
		return s.Initiated
	}
	if s.Replied != nil {
		return s.Replied
	}
	return nil
}

// GreetingEventType_Replied Proto: GreetingEventType_Replied
type GreetingEventType_Replied struct {
	ReplyMessage string `json:"replyMessage,omitempty"`
}

// GreetingEventType_Initiated Proto: GreetingEventType_Initiated
type GreetingEventType_Initiated struct {
	Name        string     `json:"name,omitempty"`
	AppVersion  string     `json:"appVersion,omitempty"`
	TestError   *TestError `json:"testError,omitempty"`
	WorkerError *TestError `json:"workerError,omitempty"`
}

// GreetingStatus Proto Enum: test.v1.GreetingStatus
type GreetingStatus string

const (
	GreetingStatus_UNSPECIFIED GreetingStatus = "UNSPECIFIED"
	GreetingStatus_INITIATED   GreetingStatus = "INITIATED"
	GreetingStatus_REPLIED     GreetingStatus = "REPLIED"
)

// GreetingKeys Proto: GreetingKeys
type GreetingKeys struct {
	GreetingId string `json:"greetingId"`
}

// GreetingState Proto: GreetingState
type GreetingState struct {
	Metadata   *state.StateMetadata `json:"metadata"`
	GreetingId string               `json:"greetingId"`
	Data       *GreetingData        `json:"data"`
	Status     GreetingStatus       `json:"status"`
}

// CombinedClient
type CombinedClient struct {
	*GreetingCommandService
	*GreetingQueryService
}

func NewCombinedClient(requester Requester) *CombinedClient {
	return &CombinedClient{
		GreetingCommandService: NewGreetingCommandService(requester),
		GreetingQueryService:   NewGreetingQueryService(requester),
	}
}
