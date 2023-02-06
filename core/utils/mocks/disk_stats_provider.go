// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	utils "github.com/smartcontractkit/chainlink/core/utils"
	mock "github.com/stretchr/testify/mock"
)

// DiskStatsProvider is an autogenerated mock type for the DiskStatsProvider type
type DiskStatsProvider struct {
	mock.Mock
}

// AvailableSpace provides a mock function with given fields: path
func (_m *DiskStatsProvider) AvailableSpace(path string) (utils.FileSize, error) {
	ret := _m.Called(path)

	var r0 utils.FileSize
	if rf, ok := ret.Get(0).(func(string) utils.FileSize); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(utils.FileSize)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDiskStatsProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewDiskStatsProvider creates a new instance of DiskStatsProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDiskStatsProvider(t mockConstructorTestingTNewDiskStatsProvider) *DiskStatsProvider {
	mock := &DiskStatsProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
