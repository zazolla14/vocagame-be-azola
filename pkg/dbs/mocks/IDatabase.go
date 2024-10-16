// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"
	dbs "vocagame-be-azola/pkg/dbs"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// IDatabase is an autogenerated mock type for the IDatabase type
type IDatabase struct {
	mock.Mock
}

// AutoMigrate provides a mock function with given fields: models
func (_m *IDatabase) AutoMigrate(models ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, models...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(models...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: ctx, model, total, opts
func (_m *IDatabase) Count(ctx context.Context, model interface{}, total *int64, opts ...dbs.FindOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, model, total)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *int64, ...dbs.FindOption) error); ok {
		r0 = rf(ctx, model, total, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, doc
func (_m *IDatabase) Create(ctx context.Context, doc interface{}) error {
	ret := _m.Called(ctx, doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInBatches provides a mock function with given fields: ctx, docs, batchSize
func (_m *IDatabase) CreateInBatches(ctx context.Context, docs interface{}, batchSize int) error {
	ret := _m.Called(ctx, docs, batchSize)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, int) error); ok {
		r0 = rf(ctx, docs, batchSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, value, opts
func (_m *IDatabase) Delete(ctx context.Context, value interface{}, opts ...dbs.FindOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...dbs.FindOption) error); ok {
		r0 = rf(ctx, value, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, result, opts
func (_m *IDatabase) Find(ctx context.Context, result interface{}, opts ...dbs.FindOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, result)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...dbs.FindOption) error); ok {
		r0 = rf(ctx, result, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: ctx, id, result
func (_m *IDatabase) FindById(ctx context.Context, id string, result interface{}) error {
	ret := _m.Called(ctx, id, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, id, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOne provides a mock function with given fields: ctx, result, opts
func (_m *IDatabase) FindOne(ctx context.Context, result interface{}, opts ...dbs.FindOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, result)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...dbs.FindOption) error); ok {
		r0 = rf(ctx, result, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDB provides a mock function with given fields:
func (_m *IDatabase) GetDB() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Update provides a mock function with given fields: ctx, doc
func (_m *IDatabase) Update(ctx context.Context, doc interface{}) error {
	ret := _m.Called(ctx, doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithTransaction provides a mock function with given fields: function
func (_m *IDatabase) WithTransaction(function func() error) error {
	ret := _m.Called(function)

	var r0 error
	if rf, ok := ret.Get(0).(func(func() error) error); ok {
		r0 = rf(function)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIDatabase creates a new instance of IDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDatabase {
	mock := &IDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
