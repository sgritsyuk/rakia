package store

import (
	"api-service/internal/domain"
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockPostStore is a mock of PostStore interface
type MockPostStore struct {
	ctrl     *gomock.Controller
	recorder *MockPostStoreMockRecorder
}

// MockPostStoreMockRecorder is the mock recorder for MockPostStore
type MockPostStoreMockRecorder struct {
	mock *MockPostStore
}

// NewMockPostStore creates a new mock instance
func NewMockPostStore(ctrl *gomock.Controller) *MockPostStore {
	mock := &MockPostStore{ctrl: ctrl}
	mock.recorder = &MockPostStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostStore) EXPECT() *MockPostStoreMockRecorder {
	return m.recorder
}

func (m *MockPostStore) Get(ctx context.Context, title string, page int, limit int) (*[]domain.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, title, page, limit)
	ret0, _ := ret[0].(*[]domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPostStoreMockRecorder) Get(ctx interface{}, title string, page int, limit int) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostStore)(nil).Get), ctx, title, page, limit)
}

func (m *MockPostStore) GetOne(ctx context.Context, id int) (*domain.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", ctx, id)
	ret0, _ := ret[0].(*domain.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPostStoreMockRecorder) GetOne(ctx interface{}, id int) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockPostStore)(nil).GetOne), ctx, id)
}

func (m *MockPostStore) Insert(ctx context.Context, post domain.Post) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, post)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPostStoreMockRecorder) Insert(ctx, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPostStore)(nil).Insert), ctx, post)
}

func (m *MockPostStore) Update(ctx context.Context, id int, post domain.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, post)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPostStoreMockRecorder) Update(ctx interface{}, id int, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostStore)(nil).Update), ctx, id, post)
}

func (m *MockPostStore) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPostStoreMockRecorder) Delete(ctx interface{}, id int) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostStore)(nil).Delete), ctx, id)
}
