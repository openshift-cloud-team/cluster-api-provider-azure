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
// Source: ../privateendpoints.go
//
// Generated by this command:
//
//	mockgen -destination privateendpoints_mock.go -package mock_privateendpoints -source ../privateendpoints.go PrivateEndpointScope
//

// Package mock_privateendpoints is a generated GoMock package.
package mock_privateendpoints

import (
	reflect "reflect"
	time "time"

	v1api20220701 "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	gomock "go.uber.org/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockPrivateEndpointScope is a mock of PrivateEndpointScope interface.
type MockPrivateEndpointScope struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointScopeMockRecorder
}

// MockPrivateEndpointScopeMockRecorder is the mock recorder for MockPrivateEndpointScope.
type MockPrivateEndpointScopeMockRecorder struct {
	mock *MockPrivateEndpointScope
}

// NewMockPrivateEndpointScope creates a new mock instance.
func NewMockPrivateEndpointScope(ctrl *gomock.Controller) *MockPrivateEndpointScope {
	mock := &MockPrivateEndpointScope{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointScope) EXPECT() *MockPrivateEndpointScopeMockRecorder {
	return m.recorder
}

// ASOOwner mocks base method.
func (m *MockPrivateEndpointScope) ASOOwner() client.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ASOOwner")
	ret0, _ := ret[0].(client.Object)
	return ret0
}

// ASOOwner indicates an expected call of ASOOwner.
func (mr *MockPrivateEndpointScopeMockRecorder) ASOOwner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ASOOwner", reflect.TypeOf((*MockPrivateEndpointScope)(nil).ASOOwner))
}

// ClusterName mocks base method.
func (m *MockPrivateEndpointScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockPrivateEndpointScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockPrivateEndpointScope)(nil).ClusterName))
}

// DefaultedAzureCallTimeout mocks base method.
func (m *MockPrivateEndpointScope) DefaultedAzureCallTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureCallTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureCallTimeout indicates an expected call of DefaultedAzureCallTimeout.
func (mr *MockPrivateEndpointScopeMockRecorder) DefaultedAzureCallTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureCallTimeout", reflect.TypeOf((*MockPrivateEndpointScope)(nil).DefaultedAzureCallTimeout))
}

// DefaultedAzureServiceReconcileTimeout mocks base method.
func (m *MockPrivateEndpointScope) DefaultedAzureServiceReconcileTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureServiceReconcileTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureServiceReconcileTimeout indicates an expected call of DefaultedAzureServiceReconcileTimeout.
func (mr *MockPrivateEndpointScopeMockRecorder) DefaultedAzureServiceReconcileTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureServiceReconcileTimeout", reflect.TypeOf((*MockPrivateEndpointScope)(nil).DefaultedAzureServiceReconcileTimeout))
}

// DefaultedReconcilerRequeue mocks base method.
func (m *MockPrivateEndpointScope) DefaultedReconcilerRequeue() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedReconcilerRequeue")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedReconcilerRequeue indicates an expected call of DefaultedReconcilerRequeue.
func (mr *MockPrivateEndpointScopeMockRecorder) DefaultedReconcilerRequeue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedReconcilerRequeue", reflect.TypeOf((*MockPrivateEndpointScope)(nil).DefaultedReconcilerRequeue))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockPrivateEndpointScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockPrivateEndpointScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockPrivateEndpointScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// GetClient mocks base method.
func (m *MockPrivateEndpointScope) GetClient() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockPrivateEndpointScopeMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockPrivateEndpointScope)(nil).GetClient))
}

// GetLongRunningOperationState mocks base method.
func (m *MockPrivateEndpointScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockPrivateEndpointScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockPrivateEndpointScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// PrivateEndpointSpecs mocks base method.
func (m *MockPrivateEndpointScope) PrivateEndpointSpecs() []azure.ASOResourceSpecGetter[*v1api20220701.PrivateEndpoint] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpointSpecs")
	ret0, _ := ret[0].([]azure.ASOResourceSpecGetter[*v1api20220701.PrivateEndpoint])
	return ret0
}

// PrivateEndpointSpecs indicates an expected call of PrivateEndpointSpecs.
func (mr *MockPrivateEndpointScopeMockRecorder) PrivateEndpointSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpointSpecs", reflect.TypeOf((*MockPrivateEndpointScope)(nil).PrivateEndpointSpecs))
}

// SetLongRunningOperationState mocks base method.
func (m *MockPrivateEndpointScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockPrivateEndpointScopeMockRecorder) SetLongRunningOperationState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockPrivateEndpointScope)(nil).SetLongRunningOperationState), arg0)
}

// UpdateDeleteStatus mocks base method.
func (m *MockPrivateEndpointScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockPrivateEndpointScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockPrivateEndpointScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockPrivateEndpointScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockPrivateEndpointScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockPrivateEndpointScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockPrivateEndpointScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockPrivateEndpointScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockPrivateEndpointScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}