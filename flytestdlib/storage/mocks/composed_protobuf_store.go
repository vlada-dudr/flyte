// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	protoiface "google.golang.org/protobuf/runtime/protoiface"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"
)

// ComposedProtobufStore is an autogenerated mock type for the ComposedProtobufStore type
type ComposedProtobufStore struct {
	mock.Mock
}

type ComposedProtobufStore_Expecter struct {
	mock *mock.Mock
}

func (_m *ComposedProtobufStore) EXPECT() *ComposedProtobufStore_Expecter {
	return &ComposedProtobufStore_Expecter{mock: &_m.Mock}
}

// CopyRaw provides a mock function with given fields: ctx, source, destination, opts
func (_m *ComposedProtobufStore) CopyRaw(ctx context.Context, source storage.DataReference, destination storage.DataReference, opts storage.Options) error {
	ret := _m.Called(ctx, source, destination, opts)

	if len(ret) == 0 {
		panic("no return value specified for CopyRaw")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.DataReference, storage.Options) error); ok {
		r0 = rf(ctx, source, destination, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComposedProtobufStore_CopyRaw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CopyRaw'
type ComposedProtobufStore_CopyRaw_Call struct {
	*mock.Call
}

// CopyRaw is a helper method to define mock.On call
//   - ctx context.Context
//   - source storage.DataReference
//   - destination storage.DataReference
//   - opts storage.Options
func (_e *ComposedProtobufStore_Expecter) CopyRaw(ctx interface{}, source interface{}, destination interface{}, opts interface{}) *ComposedProtobufStore_CopyRaw_Call {
	return &ComposedProtobufStore_CopyRaw_Call{Call: _e.mock.On("CopyRaw", ctx, source, destination, opts)}
}

func (_c *ComposedProtobufStore_CopyRaw_Call) Run(run func(ctx context.Context, source storage.DataReference, destination storage.DataReference, opts storage.Options)) *ComposedProtobufStore_CopyRaw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference), args[2].(storage.DataReference), args[3].(storage.Options))
	})
	return _c
}

func (_c *ComposedProtobufStore_CopyRaw_Call) Return(_a0 error) *ComposedProtobufStore_CopyRaw_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ComposedProtobufStore_CopyRaw_Call) RunAndReturn(run func(context.Context, storage.DataReference, storage.DataReference, storage.Options) error) *ComposedProtobufStore_CopyRaw_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSignedURL provides a mock function with given fields: ctx, reference, properties
func (_m *ComposedProtobufStore) CreateSignedURL(ctx context.Context, reference storage.DataReference, properties storage.SignedURLProperties) (storage.SignedURLResponse, error) {
	ret := _m.Called(ctx, reference, properties)

	if len(ret) == 0 {
		panic("no return value specified for CreateSignedURL")
	}

	var r0 storage.SignedURLResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.SignedURLProperties) (storage.SignedURLResponse, error)); ok {
		return rf(ctx, reference, properties)
	}
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.SignedURLProperties) storage.SignedURLResponse); ok {
		r0 = rf(ctx, reference, properties)
	} else {
		r0 = ret.Get(0).(storage.SignedURLResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference, storage.SignedURLProperties) error); ok {
		r1 = rf(ctx, reference, properties)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComposedProtobufStore_CreateSignedURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSignedURL'
type ComposedProtobufStore_CreateSignedURL_Call struct {
	*mock.Call
}

// CreateSignedURL is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
//   - properties storage.SignedURLProperties
func (_e *ComposedProtobufStore_Expecter) CreateSignedURL(ctx interface{}, reference interface{}, properties interface{}) *ComposedProtobufStore_CreateSignedURL_Call {
	return &ComposedProtobufStore_CreateSignedURL_Call{Call: _e.mock.On("CreateSignedURL", ctx, reference, properties)}
}

func (_c *ComposedProtobufStore_CreateSignedURL_Call) Run(run func(ctx context.Context, reference storage.DataReference, properties storage.SignedURLProperties)) *ComposedProtobufStore_CreateSignedURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference), args[2].(storage.SignedURLProperties))
	})
	return _c
}

func (_c *ComposedProtobufStore_CreateSignedURL_Call) Return(_a0 storage.SignedURLResponse, _a1 error) *ComposedProtobufStore_CreateSignedURL_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ComposedProtobufStore_CreateSignedURL_Call) RunAndReturn(run func(context.Context, storage.DataReference, storage.SignedURLProperties) (storage.SignedURLResponse, error)) *ComposedProtobufStore_CreateSignedURL_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, reference
func (_m *ComposedProtobufStore) Delete(ctx context.Context, reference storage.DataReference) error {
	ret := _m.Called(ctx, reference)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) error); ok {
		r0 = rf(ctx, reference)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComposedProtobufStore_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ComposedProtobufStore_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
func (_e *ComposedProtobufStore_Expecter) Delete(ctx interface{}, reference interface{}) *ComposedProtobufStore_Delete_Call {
	return &ComposedProtobufStore_Delete_Call{Call: _e.mock.On("Delete", ctx, reference)}
}

func (_c *ComposedProtobufStore_Delete_Call) Run(run func(ctx context.Context, reference storage.DataReference)) *ComposedProtobufStore_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference))
	})
	return _c
}

func (_c *ComposedProtobufStore_Delete_Call) Return(_a0 error) *ComposedProtobufStore_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ComposedProtobufStore_Delete_Call) RunAndReturn(run func(context.Context, storage.DataReference) error) *ComposedProtobufStore_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetBaseContainerFQN provides a mock function with given fields: ctx
func (_m *ComposedProtobufStore) GetBaseContainerFQN(ctx context.Context) storage.DataReference {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetBaseContainerFQN")
	}

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context) storage.DataReference); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// ComposedProtobufStore_GetBaseContainerFQN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBaseContainerFQN'
type ComposedProtobufStore_GetBaseContainerFQN_Call struct {
	*mock.Call
}

// GetBaseContainerFQN is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ComposedProtobufStore_Expecter) GetBaseContainerFQN(ctx interface{}) *ComposedProtobufStore_GetBaseContainerFQN_Call {
	return &ComposedProtobufStore_GetBaseContainerFQN_Call{Call: _e.mock.On("GetBaseContainerFQN", ctx)}
}

func (_c *ComposedProtobufStore_GetBaseContainerFQN_Call) Run(run func(ctx context.Context)) *ComposedProtobufStore_GetBaseContainerFQN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ComposedProtobufStore_GetBaseContainerFQN_Call) Return(_a0 storage.DataReference) *ComposedProtobufStore_GetBaseContainerFQN_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ComposedProtobufStore_GetBaseContainerFQN_Call) RunAndReturn(run func(context.Context) storage.DataReference) *ComposedProtobufStore_GetBaseContainerFQN_Call {
	_c.Call.Return(run)
	return _c
}

// Head provides a mock function with given fields: ctx, reference
func (_m *ComposedProtobufStore) Head(ctx context.Context, reference storage.DataReference) (storage.Metadata, error) {
	ret := _m.Called(ctx, reference)

	if len(ret) == 0 {
		panic("no return value specified for Head")
	}

	var r0 storage.Metadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) (storage.Metadata, error)); ok {
		return rf(ctx, reference)
	}
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) storage.Metadata); ok {
		r0 = rf(ctx, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Metadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference) error); ok {
		r1 = rf(ctx, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComposedProtobufStore_Head_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Head'
type ComposedProtobufStore_Head_Call struct {
	*mock.Call
}

// Head is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
func (_e *ComposedProtobufStore_Expecter) Head(ctx interface{}, reference interface{}) *ComposedProtobufStore_Head_Call {
	return &ComposedProtobufStore_Head_Call{Call: _e.mock.On("Head", ctx, reference)}
}

func (_c *ComposedProtobufStore_Head_Call) Run(run func(ctx context.Context, reference storage.DataReference)) *ComposedProtobufStore_Head_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference))
	})
	return _c
}

func (_c *ComposedProtobufStore_Head_Call) Return(_a0 storage.Metadata, _a1 error) *ComposedProtobufStore_Head_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ComposedProtobufStore_Head_Call) RunAndReturn(run func(context.Context, storage.DataReference) (storage.Metadata, error)) *ComposedProtobufStore_Head_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, reference, maxItems, cursor
func (_m *ComposedProtobufStore) List(ctx context.Context, reference storage.DataReference, maxItems int, cursor storage.Cursor) ([]storage.DataReference, storage.Cursor, error) {
	ret := _m.Called(ctx, reference, maxItems, cursor)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []storage.DataReference
	var r1 storage.Cursor
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, int, storage.Cursor) ([]storage.DataReference, storage.Cursor, error)); ok {
		return rf(ctx, reference, maxItems, cursor)
	}
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, int, storage.Cursor) []storage.DataReference); ok {
		r0 = rf(ctx, reference, maxItems, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.DataReference)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference, int, storage.Cursor) storage.Cursor); ok {
		r1 = rf(ctx, reference, maxItems, cursor)
	} else {
		r1 = ret.Get(1).(storage.Cursor)
	}

	if rf, ok := ret.Get(2).(func(context.Context, storage.DataReference, int, storage.Cursor) error); ok {
		r2 = rf(ctx, reference, maxItems, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ComposedProtobufStore_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ComposedProtobufStore_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
//   - maxItems int
//   - cursor storage.Cursor
func (_e *ComposedProtobufStore_Expecter) List(ctx interface{}, reference interface{}, maxItems interface{}, cursor interface{}) *ComposedProtobufStore_List_Call {
	return &ComposedProtobufStore_List_Call{Call: _e.mock.On("List", ctx, reference, maxItems, cursor)}
}

func (_c *ComposedProtobufStore_List_Call) Run(run func(ctx context.Context, reference storage.DataReference, maxItems int, cursor storage.Cursor)) *ComposedProtobufStore_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference), args[2].(int), args[3].(storage.Cursor))
	})
	return _c
}

func (_c *ComposedProtobufStore_List_Call) Return(_a0 []storage.DataReference, _a1 storage.Cursor, _a2 error) *ComposedProtobufStore_List_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ComposedProtobufStore_List_Call) RunAndReturn(run func(context.Context, storage.DataReference, int, storage.Cursor) ([]storage.DataReference, storage.Cursor, error)) *ComposedProtobufStore_List_Call {
	_c.Call.Return(run)
	return _c
}

// ReadProtobuf provides a mock function with given fields: ctx, reference, msg
func (_m *ComposedProtobufStore) ReadProtobuf(ctx context.Context, reference storage.DataReference, msg protoiface.MessageV1) error {
	ret := _m.Called(ctx, reference, msg)

	if len(ret) == 0 {
		panic("no return value specified for ReadProtobuf")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, protoiface.MessageV1) error); ok {
		r0 = rf(ctx, reference, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComposedProtobufStore_ReadProtobuf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadProtobuf'
type ComposedProtobufStore_ReadProtobuf_Call struct {
	*mock.Call
}

// ReadProtobuf is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
//   - msg protoiface.MessageV1
func (_e *ComposedProtobufStore_Expecter) ReadProtobuf(ctx interface{}, reference interface{}, msg interface{}) *ComposedProtobufStore_ReadProtobuf_Call {
	return &ComposedProtobufStore_ReadProtobuf_Call{Call: _e.mock.On("ReadProtobuf", ctx, reference, msg)}
}

func (_c *ComposedProtobufStore_ReadProtobuf_Call) Run(run func(ctx context.Context, reference storage.DataReference, msg protoiface.MessageV1)) *ComposedProtobufStore_ReadProtobuf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference), args[2].(protoiface.MessageV1))
	})
	return _c
}

func (_c *ComposedProtobufStore_ReadProtobuf_Call) Return(_a0 error) *ComposedProtobufStore_ReadProtobuf_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ComposedProtobufStore_ReadProtobuf_Call) RunAndReturn(run func(context.Context, storage.DataReference, protoiface.MessageV1) error) *ComposedProtobufStore_ReadProtobuf_Call {
	_c.Call.Return(run)
	return _c
}

// ReadRaw provides a mock function with given fields: ctx, reference
func (_m *ComposedProtobufStore) ReadRaw(ctx context.Context, reference storage.DataReference) (io.ReadCloser, error) {
	ret := _m.Called(ctx, reference)

	if len(ret) == 0 {
		panic("no return value specified for ReadRaw")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) (io.ReadCloser, error)); ok {
		return rf(ctx, reference)
	}
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) io.ReadCloser); ok {
		r0 = rf(ctx, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference) error); ok {
		r1 = rf(ctx, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComposedProtobufStore_ReadRaw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadRaw'
type ComposedProtobufStore_ReadRaw_Call struct {
	*mock.Call
}

// ReadRaw is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
func (_e *ComposedProtobufStore_Expecter) ReadRaw(ctx interface{}, reference interface{}) *ComposedProtobufStore_ReadRaw_Call {
	return &ComposedProtobufStore_ReadRaw_Call{Call: _e.mock.On("ReadRaw", ctx, reference)}
}

func (_c *ComposedProtobufStore_ReadRaw_Call) Run(run func(ctx context.Context, reference storage.DataReference)) *ComposedProtobufStore_ReadRaw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference))
	})
	return _c
}

func (_c *ComposedProtobufStore_ReadRaw_Call) Return(_a0 io.ReadCloser, _a1 error) *ComposedProtobufStore_ReadRaw_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ComposedProtobufStore_ReadRaw_Call) RunAndReturn(run func(context.Context, storage.DataReference) (io.ReadCloser, error)) *ComposedProtobufStore_ReadRaw_Call {
	_c.Call.Return(run)
	return _c
}

// WriteProtobuf provides a mock function with given fields: ctx, reference, opts, msg
func (_m *ComposedProtobufStore) WriteProtobuf(ctx context.Context, reference storage.DataReference, opts storage.Options, msg protoiface.MessageV1) error {
	ret := _m.Called(ctx, reference, opts, msg)

	if len(ret) == 0 {
		panic("no return value specified for WriteProtobuf")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.Options, protoiface.MessageV1) error); ok {
		r0 = rf(ctx, reference, opts, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComposedProtobufStore_WriteProtobuf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteProtobuf'
type ComposedProtobufStore_WriteProtobuf_Call struct {
	*mock.Call
}

// WriteProtobuf is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
//   - opts storage.Options
//   - msg protoiface.MessageV1
func (_e *ComposedProtobufStore_Expecter) WriteProtobuf(ctx interface{}, reference interface{}, opts interface{}, msg interface{}) *ComposedProtobufStore_WriteProtobuf_Call {
	return &ComposedProtobufStore_WriteProtobuf_Call{Call: _e.mock.On("WriteProtobuf", ctx, reference, opts, msg)}
}

func (_c *ComposedProtobufStore_WriteProtobuf_Call) Run(run func(ctx context.Context, reference storage.DataReference, opts storage.Options, msg protoiface.MessageV1)) *ComposedProtobufStore_WriteProtobuf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference), args[2].(storage.Options), args[3].(protoiface.MessageV1))
	})
	return _c
}

func (_c *ComposedProtobufStore_WriteProtobuf_Call) Return(_a0 error) *ComposedProtobufStore_WriteProtobuf_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ComposedProtobufStore_WriteProtobuf_Call) RunAndReturn(run func(context.Context, storage.DataReference, storage.Options, protoiface.MessageV1) error) *ComposedProtobufStore_WriteProtobuf_Call {
	_c.Call.Return(run)
	return _c
}

// WriteRaw provides a mock function with given fields: ctx, reference, size, opts, raw
func (_m *ComposedProtobufStore) WriteRaw(ctx context.Context, reference storage.DataReference, size int64, opts storage.Options, raw io.Reader) error {
	ret := _m.Called(ctx, reference, size, opts, raw)

	if len(ret) == 0 {
		panic("no return value specified for WriteRaw")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, int64, storage.Options, io.Reader) error); ok {
		r0 = rf(ctx, reference, size, opts, raw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComposedProtobufStore_WriteRaw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteRaw'
type ComposedProtobufStore_WriteRaw_Call struct {
	*mock.Call
}

// WriteRaw is a helper method to define mock.On call
//   - ctx context.Context
//   - reference storage.DataReference
//   - size int64
//   - opts storage.Options
//   - raw io.Reader
func (_e *ComposedProtobufStore_Expecter) WriteRaw(ctx interface{}, reference interface{}, size interface{}, opts interface{}, raw interface{}) *ComposedProtobufStore_WriteRaw_Call {
	return &ComposedProtobufStore_WriteRaw_Call{Call: _e.mock.On("WriteRaw", ctx, reference, size, opts, raw)}
}

func (_c *ComposedProtobufStore_WriteRaw_Call) Run(run func(ctx context.Context, reference storage.DataReference, size int64, opts storage.Options, raw io.Reader)) *ComposedProtobufStore_WriteRaw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(storage.DataReference), args[2].(int64), args[3].(storage.Options), args[4].(io.Reader))
	})
	return _c
}

func (_c *ComposedProtobufStore_WriteRaw_Call) Return(_a0 error) *ComposedProtobufStore_WriteRaw_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ComposedProtobufStore_WriteRaw_Call) RunAndReturn(run func(context.Context, storage.DataReference, int64, storage.Options, io.Reader) error) *ComposedProtobufStore_WriteRaw_Call {
	_c.Call.Return(run)
	return _c
}

// NewComposedProtobufStore creates a new instance of ComposedProtobufStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComposedProtobufStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *ComposedProtobufStore {
	mock := &ComposedProtobufStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
