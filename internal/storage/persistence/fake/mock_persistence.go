// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/storage/persistence (interfaces: Storage)
//
// Generated by this command:
//
//	mockgen -typed -destination=./internal/storage/persistence/fake/mock_persistence.go -package=fake github.com/artefactual-sdps/enduro/internal/storage/persistence Storage
//

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	types "github.com/artefactual-sdps/enduro/internal/storage/types"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateLocation mocks base method.
func (m *MockStorage) CreateLocation(arg0 context.Context, arg1 *storage.Location, arg2 *types.LocationConfig) (*storage.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation.
func (mr *MockStorageMockRecorder) CreateLocation(arg0, arg1, arg2 any) *MockStorageCreateLocationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockStorage)(nil).CreateLocation), arg0, arg1, arg2)
	return &MockStorageCreateLocationCall{Call: call}
}

// MockStorageCreateLocationCall wrap *gomock.Call
type MockStorageCreateLocationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageCreateLocationCall) Return(arg0 *storage.Location, arg1 error) *MockStorageCreateLocationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageCreateLocationCall) Do(f func(context.Context, *storage.Location, *types.LocationConfig) (*storage.Location, error)) *MockStorageCreateLocationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageCreateLocationCall) DoAndReturn(f func(context.Context, *storage.Location, *types.LocationConfig) (*storage.Location, error)) *MockStorageCreateLocationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreatePackage mocks base method.
func (m *MockStorage) CreatePackage(arg0 context.Context, arg1 *storage.Package) (*storage.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePackage", arg0, arg1)
	ret0, _ := ret[0].(*storage.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePackage indicates an expected call of CreatePackage.
func (mr *MockStorageMockRecorder) CreatePackage(arg0, arg1 any) *MockStorageCreatePackageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePackage", reflect.TypeOf((*MockStorage)(nil).CreatePackage), arg0, arg1)
	return &MockStorageCreatePackageCall{Call: call}
}

// MockStorageCreatePackageCall wrap *gomock.Call
type MockStorageCreatePackageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageCreatePackageCall) Return(arg0 *storage.Package, arg1 error) *MockStorageCreatePackageCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageCreatePackageCall) Do(f func(context.Context, *storage.Package) (*storage.Package, error)) *MockStorageCreatePackageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageCreatePackageCall) DoAndReturn(f func(context.Context, *storage.Package) (*storage.Package, error)) *MockStorageCreatePackageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListLocations mocks base method.
func (m *MockStorage) ListLocations(arg0 context.Context) (storage.LocationCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", arg0)
	ret0, _ := ret[0].(storage.LocationCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockStorageMockRecorder) ListLocations(arg0 any) *MockStorageListLocationsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockStorage)(nil).ListLocations), arg0)
	return &MockStorageListLocationsCall{Call: call}
}

// MockStorageListLocationsCall wrap *gomock.Call
type MockStorageListLocationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageListLocationsCall) Return(arg0 storage.LocationCollection, arg1 error) *MockStorageListLocationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageListLocationsCall) Do(f func(context.Context) (storage.LocationCollection, error)) *MockStorageListLocationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageListLocationsCall) DoAndReturn(f func(context.Context) (storage.LocationCollection, error)) *MockStorageListLocationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListPackages mocks base method.
func (m *MockStorage) ListPackages(arg0 context.Context) (storage.PackageCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPackages", arg0)
	ret0, _ := ret[0].(storage.PackageCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPackages indicates an expected call of ListPackages.
func (mr *MockStorageMockRecorder) ListPackages(arg0 any) *MockStorageListPackagesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPackages", reflect.TypeOf((*MockStorage)(nil).ListPackages), arg0)
	return &MockStorageListPackagesCall{Call: call}
}

// MockStorageListPackagesCall wrap *gomock.Call
type MockStorageListPackagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageListPackagesCall) Return(arg0 storage.PackageCollection, arg1 error) *MockStorageListPackagesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageListPackagesCall) Do(f func(context.Context) (storage.PackageCollection, error)) *MockStorageListPackagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageListPackagesCall) DoAndReturn(f func(context.Context) (storage.PackageCollection, error)) *MockStorageListPackagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LocationPackages mocks base method.
func (m *MockStorage) LocationPackages(arg0 context.Context, arg1 uuid.UUID) (storage.PackageCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationPackages", arg0, arg1)
	ret0, _ := ret[0].(storage.PackageCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationPackages indicates an expected call of LocationPackages.
func (mr *MockStorageMockRecorder) LocationPackages(arg0, arg1 any) *MockStorageLocationPackagesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationPackages", reflect.TypeOf((*MockStorage)(nil).LocationPackages), arg0, arg1)
	return &MockStorageLocationPackagesCall{Call: call}
}

// MockStorageLocationPackagesCall wrap *gomock.Call
type MockStorageLocationPackagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageLocationPackagesCall) Return(arg0 storage.PackageCollection, arg1 error) *MockStorageLocationPackagesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageLocationPackagesCall) Do(f func(context.Context, uuid.UUID) (storage.PackageCollection, error)) *MockStorageLocationPackagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageLocationPackagesCall) DoAndReturn(f func(context.Context, uuid.UUID) (storage.PackageCollection, error)) *MockStorageLocationPackagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadLocation mocks base method.
func (m *MockStorage) ReadLocation(arg0 context.Context, arg1 uuid.UUID) (*storage.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLocation", arg0, arg1)
	ret0, _ := ret[0].(*storage.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadLocation indicates an expected call of ReadLocation.
func (mr *MockStorageMockRecorder) ReadLocation(arg0, arg1 any) *MockStorageReadLocationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLocation", reflect.TypeOf((*MockStorage)(nil).ReadLocation), arg0, arg1)
	return &MockStorageReadLocationCall{Call: call}
}

// MockStorageReadLocationCall wrap *gomock.Call
type MockStorageReadLocationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageReadLocationCall) Return(arg0 *storage.Location, arg1 error) *MockStorageReadLocationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageReadLocationCall) Do(f func(context.Context, uuid.UUID) (*storage.Location, error)) *MockStorageReadLocationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageReadLocationCall) DoAndReturn(f func(context.Context, uuid.UUID) (*storage.Location, error)) *MockStorageReadLocationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadPackage mocks base method.
func (m *MockStorage) ReadPackage(arg0 context.Context, arg1 uuid.UUID) (*storage.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPackage", arg0, arg1)
	ret0, _ := ret[0].(*storage.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPackage indicates an expected call of ReadPackage.
func (mr *MockStorageMockRecorder) ReadPackage(arg0, arg1 any) *MockStorageReadPackageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPackage", reflect.TypeOf((*MockStorage)(nil).ReadPackage), arg0, arg1)
	return &MockStorageReadPackageCall{Call: call}
}

// MockStorageReadPackageCall wrap *gomock.Call
type MockStorageReadPackageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageReadPackageCall) Return(arg0 *storage.Package, arg1 error) *MockStorageReadPackageCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageReadPackageCall) Do(f func(context.Context, uuid.UUID) (*storage.Package, error)) *MockStorageReadPackageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageReadPackageCall) DoAndReturn(f func(context.Context, uuid.UUID) (*storage.Package, error)) *MockStorageReadPackageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePackageLocationID mocks base method.
func (m *MockStorage) UpdatePackageLocationID(arg0 context.Context, arg1, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageLocationID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageLocationID indicates an expected call of UpdatePackageLocationID.
func (mr *MockStorageMockRecorder) UpdatePackageLocationID(arg0, arg1, arg2 any) *MockStorageUpdatePackageLocationIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageLocationID", reflect.TypeOf((*MockStorage)(nil).UpdatePackageLocationID), arg0, arg1, arg2)
	return &MockStorageUpdatePackageLocationIDCall{Call: call}
}

// MockStorageUpdatePackageLocationIDCall wrap *gomock.Call
type MockStorageUpdatePackageLocationIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageUpdatePackageLocationIDCall) Return(arg0 error) *MockStorageUpdatePackageLocationIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageUpdatePackageLocationIDCall) Do(f func(context.Context, uuid.UUID, uuid.UUID) error) *MockStorageUpdatePackageLocationIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageUpdatePackageLocationIDCall) DoAndReturn(f func(context.Context, uuid.UUID, uuid.UUID) error) *MockStorageUpdatePackageLocationIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdatePackageStatus mocks base method.
func (m *MockStorage) UpdatePackageStatus(arg0 context.Context, arg1 uuid.UUID, arg2 types.PackageStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackageStatus indicates an expected call of UpdatePackageStatus.
func (mr *MockStorageMockRecorder) UpdatePackageStatus(arg0, arg1, arg2 any) *MockStorageUpdatePackageStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageStatus", reflect.TypeOf((*MockStorage)(nil).UpdatePackageStatus), arg0, arg1, arg2)
	return &MockStorageUpdatePackageStatusCall{Call: call}
}

// MockStorageUpdatePackageStatusCall wrap *gomock.Call
type MockStorageUpdatePackageStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageUpdatePackageStatusCall) Return(arg0 error) *MockStorageUpdatePackageStatusCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageUpdatePackageStatusCall) Do(f func(context.Context, uuid.UUID, types.PackageStatus) error) *MockStorageUpdatePackageStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageUpdatePackageStatusCall) DoAndReturn(f func(context.Context, uuid.UUID, types.PackageStatus) error) *MockStorageUpdatePackageStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
