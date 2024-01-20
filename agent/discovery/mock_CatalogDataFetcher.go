// Code generated by mockery v2.32.4. DO NOT EDIT.

package discovery

import (
	config "github.com/hashicorp/consul/agent/config"
	mock "github.com/stretchr/testify/mock"

	net "net"
)

// MockCatalogDataFetcher is an autogenerated mock type for the CatalogDataFetcher type
type MockCatalogDataFetcher struct {
	mock.Mock
}

// FetchEndpoints provides a mock function with given fields: ctx, req, lookupType
func (_m *MockCatalogDataFetcher) FetchEndpoints(ctx Context, req *QueryPayload, lookupType LookupType) ([]*Result, error) {
	ret := _m.Called(ctx, req, lookupType)

	var r0 []*Result
	var r1 error
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload, LookupType) ([]*Result, error)); ok {
		return rf(ctx, req, lookupType)
	}
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload, LookupType) []*Result); ok {
		r0 = rf(ctx, req, lookupType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Result)
		}
	}

	if rf, ok := ret.Get(1).(func(Context, *QueryPayload, LookupType) error); ok {
		r1 = rf(ctx, req, lookupType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchNodes provides a mock function with given fields: ctx, req
func (_m *MockCatalogDataFetcher) FetchNodes(ctx Context, req *QueryPayload) ([]*Result, error) {
	ret := _m.Called(ctx, req)

	var r0 []*Result
	var r1 error
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) ([]*Result, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) []*Result); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Result)
		}
	}

	if rf, ok := ret.Get(1).(func(Context, *QueryPayload) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchPreparedQuery provides a mock function with given fields: ctx, req
func (_m *MockCatalogDataFetcher) FetchPreparedQuery(ctx Context, req *QueryPayload) ([]*Result, error) {
	ret := _m.Called(ctx, req)

	var r0 []*Result
	var r1 error
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) ([]*Result, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) []*Result); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Result)
		}
	}

	if rf, ok := ret.Get(1).(func(Context, *QueryPayload) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchRecordsByIp provides a mock function with given fields: ctx, ip
func (_m *MockCatalogDataFetcher) FetchRecordsByIp(ctx Context, ip net.IP) ([]*Result, error) {
	ret := _m.Called(ctx, ip)

	var r0 []*Result
	var r1 error
	if rf, ok := ret.Get(0).(func(Context, net.IP) ([]*Result, error)); ok {
		return rf(ctx, ip)
	}
	if rf, ok := ret.Get(0).(func(Context, net.IP) []*Result); ok {
		r0 = rf(ctx, ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Result)
		}
	}

	if rf, ok := ret.Get(1).(func(Context, net.IP) error); ok {
		r1 = rf(ctx, ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchVirtualIP provides a mock function with given fields: ctx, req
func (_m *MockCatalogDataFetcher) FetchVirtualIP(ctx Context, req *QueryPayload) (*Result, error) {
	ret := _m.Called(ctx, req)

	var r0 *Result
	var r1 error
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) (*Result, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) *Result); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Result)
		}
	}

	if rf, ok := ret.Get(1).(func(Context, *QueryPayload) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchWorkload provides a mock function with given fields: ctx, req
func (_m *MockCatalogDataFetcher) FetchWorkload(ctx Context, req *QueryPayload) (*Result, error) {
	ret := _m.Called(ctx, req)

	var r0 *Result
	var r1 error
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) (*Result, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(Context, *QueryPayload) *Result); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Result)
		}
	}

	if rf, ok := ret.Get(1).(func(Context, *QueryPayload) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadConfig provides a mock function with given fields: _a0
func (_m *MockCatalogDataFetcher) LoadConfig(_a0 *config.RuntimeConfig) {
	_m.Called(_a0)
}

// NewMockCatalogDataFetcher creates a new instance of MockCatalogDataFetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCatalogDataFetcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCatalogDataFetcher {
	mock := &MockCatalogDataFetcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
