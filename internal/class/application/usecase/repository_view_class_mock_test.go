// Code generated by mockery v2.30.1. DO NOT EDIT.

package usecase_test

import (
	domain "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryViewClassMock is an autogenerated mock type for the RepositoryViewClass type
type RepositoryViewClassMock struct {
	mock.Mock
}

// GetClassByClassID provides a mock function with given fields: classID
func (_m *RepositoryViewClassMock) GetClassByClassID(classID string) (domain.Class, error) {
	ret := _m.Called(classID)

	var r0 domain.Class
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.Class, error)); ok {
		return rf(classID)
	}
	if rf, ok := ret.Get(0).(func(string) domain.Class); ok {
		r0 = rf(classID)
	} else {
		r0 = ret.Get(0).(domain.Class)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(classID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepositoryViewClassMock creates a new instance of RepositoryViewClassMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryViewClassMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryViewClassMock {
	mock := &RepositoryViewClassMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
