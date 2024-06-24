/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go
//
// Generated by this command:
//
//	mockgen -destination client_mock.go -package mock_resourcehealth -source ../client.go Client
//

// Package mock_resourcehealth is a generated GoMock package.
package mock_resourcehealth

import (
	context "context"
	reflect "reflect"

	armresourcehealth "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
	gomock "go.uber.org/mock/gomock"
)

// Mockclient is a mock of client interface.
type Mockclient struct {
	ctrl     *gomock.Controller
	recorder *MockclientMockRecorder
}

// MockclientMockRecorder is the mock recorder for Mockclient.
type MockclientMockRecorder struct {
	mock *Mockclient
}

// NewMockclient creates a new mock instance.
func NewMockclient(ctrl *gomock.Controller) *Mockclient {
	mock := &Mockclient{ctrl: ctrl}
	mock.recorder = &MockclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockclient) EXPECT() *MockclientMockRecorder {
	return m.recorder
}

// GetByResource mocks base method.
func (m *Mockclient) GetByResource(arg0 context.Context, arg1 string) (armresourcehealth.AvailabilityStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByResource", arg0, arg1)
	ret0, _ := ret[0].(armresourcehealth.AvailabilityStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByResource indicates an expected call of GetByResource.
func (mr *MockclientMockRecorder) GetByResource(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByResource", reflect.TypeOf((*Mockclient)(nil).GetByResource), arg0, arg1)
}