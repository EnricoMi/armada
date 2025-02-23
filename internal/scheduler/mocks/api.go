// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/armadaproject/armada/pkg/api (interfaces: SubmitClient,Submit_GetQueuesClient)
//
// Generated by this command:
//
//	mockgen -destination=./api.go -package=schedulermocks github.com/armadaproject/armada/pkg/api SubmitClient,Submit_GetQueuesClient
//

// Package schedulermocks is a generated GoMock package.
package schedulermocks

import (
	context "context"
	reflect "reflect"

	api "github.com/armadaproject/armada/pkg/api"
	types "github.com/gogo/protobuf/types"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockSubmitClient is a mock of SubmitClient interface.
type MockSubmitClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubmitClientMockRecorder
	isgomock struct{}
}

// MockSubmitClientMockRecorder is the mock recorder for MockSubmitClient.
type MockSubmitClientMockRecorder struct {
	mock *MockSubmitClient
}

// NewMockSubmitClient creates a new mock instance.
func NewMockSubmitClient(ctrl *gomock.Controller) *MockSubmitClient {
	mock := &MockSubmitClient{ctrl: ctrl}
	mock.recorder = &MockSubmitClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubmitClient) EXPECT() *MockSubmitClientMockRecorder {
	return m.recorder
}

// CancelJobSet mocks base method.
func (m *MockSubmitClient) CancelJobSet(ctx context.Context, in *api.JobSetCancelRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelJobSet", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelJobSet indicates an expected call of CancelJobSet.
func (mr *MockSubmitClientMockRecorder) CancelJobSet(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelJobSet", reflect.TypeOf((*MockSubmitClient)(nil).CancelJobSet), varargs...)
}

// CancelJobs mocks base method.
func (m *MockSubmitClient) CancelJobs(ctx context.Context, in *api.JobCancelRequest, opts ...grpc.CallOption) (*api.CancellationResult, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelJobs", varargs...)
	ret0, _ := ret[0].(*api.CancellationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelJobs indicates an expected call of CancelJobs.
func (mr *MockSubmitClientMockRecorder) CancelJobs(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelJobs", reflect.TypeOf((*MockSubmitClient)(nil).CancelJobs), varargs...)
}

// CreateQueue mocks base method.
func (m *MockSubmitClient) CreateQueue(ctx context.Context, in *api.Queue, opts ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateQueue", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQueue indicates an expected call of CreateQueue.
func (mr *MockSubmitClientMockRecorder) CreateQueue(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQueue", reflect.TypeOf((*MockSubmitClient)(nil).CreateQueue), varargs...)
}

// CreateQueues mocks base method.
func (m *MockSubmitClient) CreateQueues(ctx context.Context, in *api.QueueList, opts ...grpc.CallOption) (*api.BatchQueueCreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateQueues", varargs...)
	ret0, _ := ret[0].(*api.BatchQueueCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQueues indicates an expected call of CreateQueues.
func (mr *MockSubmitClientMockRecorder) CreateQueues(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQueues", reflect.TypeOf((*MockSubmitClient)(nil).CreateQueues), varargs...)
}

// DeleteQueue mocks base method.
func (m *MockSubmitClient) DeleteQueue(ctx context.Context, in *api.QueueDeleteRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteQueue", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteQueue indicates an expected call of DeleteQueue.
func (mr *MockSubmitClientMockRecorder) DeleteQueue(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQueue", reflect.TypeOf((*MockSubmitClient)(nil).DeleteQueue), varargs...)
}

// GetQueue mocks base method.
func (m *MockSubmitClient) GetQueue(ctx context.Context, in *api.QueueGetRequest, opts ...grpc.CallOption) (*api.Queue, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueue", varargs...)
	ret0, _ := ret[0].(*api.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueue indicates an expected call of GetQueue.
func (mr *MockSubmitClientMockRecorder) GetQueue(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueue", reflect.TypeOf((*MockSubmitClient)(nil).GetQueue), varargs...)
}

// GetQueues mocks base method.
func (m *MockSubmitClient) GetQueues(ctx context.Context, in *api.StreamingQueueGetRequest, opts ...grpc.CallOption) (api.Submit_GetQueuesClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueues", varargs...)
	ret0, _ := ret[0].(api.Submit_GetQueuesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueues indicates an expected call of GetQueues.
func (mr *MockSubmitClientMockRecorder) GetQueues(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueues", reflect.TypeOf((*MockSubmitClient)(nil).GetQueues), varargs...)
}

// Health mocks base method.
func (m *MockSubmitClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*api.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Health", varargs...)
	ret0, _ := ret[0].(*api.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health.
func (mr *MockSubmitClientMockRecorder) Health(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockSubmitClient)(nil).Health), varargs...)
}

// PreemptJobs mocks base method.
func (m *MockSubmitClient) PreemptJobs(ctx context.Context, in *api.JobPreemptRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreemptJobs", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreemptJobs indicates an expected call of PreemptJobs.
func (mr *MockSubmitClientMockRecorder) PreemptJobs(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreemptJobs", reflect.TypeOf((*MockSubmitClient)(nil).PreemptJobs), varargs...)
}

// ReprioritizeJobs mocks base method.
func (m *MockSubmitClient) ReprioritizeJobs(ctx context.Context, in *api.JobReprioritizeRequest, opts ...grpc.CallOption) (*api.JobReprioritizeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReprioritizeJobs", varargs...)
	ret0, _ := ret[0].(*api.JobReprioritizeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReprioritizeJobs indicates an expected call of ReprioritizeJobs.
func (mr *MockSubmitClientMockRecorder) ReprioritizeJobs(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReprioritizeJobs", reflect.TypeOf((*MockSubmitClient)(nil).ReprioritizeJobs), varargs...)
}

// SubmitJobs mocks base method.
func (m *MockSubmitClient) SubmitJobs(ctx context.Context, in *api.JobSubmitRequest, opts ...grpc.CallOption) (*api.JobSubmitResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitJobs", varargs...)
	ret0, _ := ret[0].(*api.JobSubmitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitJobs indicates an expected call of SubmitJobs.
func (mr *MockSubmitClientMockRecorder) SubmitJobs(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitJobs", reflect.TypeOf((*MockSubmitClient)(nil).SubmitJobs), varargs...)
}

// UpdateQueue mocks base method.
func (m *MockSubmitClient) UpdateQueue(ctx context.Context, in *api.Queue, opts ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateQueue", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQueue indicates an expected call of UpdateQueue.
func (mr *MockSubmitClientMockRecorder) UpdateQueue(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQueue", reflect.TypeOf((*MockSubmitClient)(nil).UpdateQueue), varargs...)
}

// UpdateQueues mocks base method.
func (m *MockSubmitClient) UpdateQueues(ctx context.Context, in *api.QueueList, opts ...grpc.CallOption) (*api.BatchQueueUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateQueues", varargs...)
	ret0, _ := ret[0].(*api.BatchQueueUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQueues indicates an expected call of UpdateQueues.
func (mr *MockSubmitClientMockRecorder) UpdateQueues(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQueues", reflect.TypeOf((*MockSubmitClient)(nil).UpdateQueues), varargs...)
}

// MockSubmit_GetQueuesClient is a mock of Submit_GetQueuesClient interface.
type MockSubmit_GetQueuesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubmit_GetQueuesClientMockRecorder
	isgomock struct{}
}

// MockSubmit_GetQueuesClientMockRecorder is the mock recorder for MockSubmit_GetQueuesClient.
type MockSubmit_GetQueuesClientMockRecorder struct {
	mock *MockSubmit_GetQueuesClient
}

// NewMockSubmit_GetQueuesClient creates a new mock instance.
func NewMockSubmit_GetQueuesClient(ctrl *gomock.Controller) *MockSubmit_GetQueuesClient {
	mock := &MockSubmit_GetQueuesClient{ctrl: ctrl}
	mock.recorder = &MockSubmit_GetQueuesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubmit_GetQueuesClient) EXPECT() *MockSubmit_GetQueuesClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockSubmit_GetQueuesClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockSubmit_GetQueuesClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockSubmit_GetQueuesClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockSubmit_GetQueuesClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSubmit_GetQueuesClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSubmit_GetQueuesClient)(nil).Context))
}

// Header mocks base method.
func (m *MockSubmit_GetQueuesClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockSubmit_GetQueuesClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockSubmit_GetQueuesClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockSubmit_GetQueuesClient) Recv() (*api.StreamingQueueMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*api.StreamingQueueMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockSubmit_GetQueuesClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockSubmit_GetQueuesClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockSubmit_GetQueuesClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockSubmit_GetQueuesClientMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSubmit_GetQueuesClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockSubmit_GetQueuesClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSubmit_GetQueuesClientMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSubmit_GetQueuesClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockSubmit_GetQueuesClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockSubmit_GetQueuesClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockSubmit_GetQueuesClient)(nil).Trailer))
}
