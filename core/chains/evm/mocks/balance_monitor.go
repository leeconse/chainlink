// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	assets "github.com/smartcontractkit/chainlink/core/assets"

	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"
)

// BalanceMonitor is an autogenerated mock type for the BalanceMonitor type
type BalanceMonitor struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *BalanceMonitor) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEthBalance provides a mock function with given fields: _a0
func (_m *BalanceMonitor) GetEthBalance(_a0 common.Address) *assets.Eth {
	ret := _m.Called(_a0)

	var r0 *assets.Eth
	if rf, ok := ret.Get(0).(func(common.Address) *assets.Eth); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Eth)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *BalanceMonitor) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *BalanceMonitor) OnNewLongestChain(ctx context.Context, head *types.Head) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *BalanceMonitor) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *BalanceMonitor) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBalanceMonitor interface {
	mock.TestingT
	Cleanup(func())
}

// NewBalanceMonitor creates a new instance of BalanceMonitor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBalanceMonitor(t mockConstructorTestingTNewBalanceMonitor) *BalanceMonitor {
	mock := &BalanceMonitor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
