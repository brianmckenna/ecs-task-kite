// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/aws-sdk-go/service/ecs/ecsiface (interfaces: ECSAPI)

package mock_ecsiface

import (
	gomock "github.com/golang/mock/gomock"
	ecs "github.com/aws/aws-sdk-go/service/ecs"
)

// Mock of ECSAPI interface
type MockECSAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockECSAPIRecorder
}

// Recorder for MockECSAPI (not exported)
type _MockECSAPIRecorder struct {
	mock *MockECSAPI
}

func NewMockECSAPI(ctrl *gomock.Controller) *MockECSAPI {
	mock := &MockECSAPI{ctrl: ctrl}
	mock.recorder = &_MockECSAPIRecorder{mock}
	return mock
}

func (_m *MockECSAPI) EXPECT() *_MockECSAPIRecorder {
	return _m.recorder
}

func (_m *MockECSAPI) CreateCluster(_param0 *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateCluster", _param0)
	ret0, _ := ret[0].(*ecs.CreateClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateCluster", arg0)
}

func (_m *MockECSAPI) CreateService(_param0 *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateService", _param0)
	ret0, _ := ret[0].(*ecs.CreateServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) CreateService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateService", arg0)
}

func (_m *MockECSAPI) DeleteCluster(_param0 *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteCluster", _param0)
	ret0, _ := ret[0].(*ecs.DeleteClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCluster", arg0)
}

func (_m *MockECSAPI) DeleteService(_param0 *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteService", _param0)
	ret0, _ := ret[0].(*ecs.DeleteServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeleteService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteService", arg0)
}

func (_m *MockECSAPI) DeregisterContainerInstance(_param0 *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error) {
	ret := _m.ctrl.Call(_m, "DeregisterContainerInstance", _param0)
	ret0, _ := ret[0].(*ecs.DeregisterContainerInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeregisterContainerInstance(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterContainerInstance", arg0)
}

func (_m *MockECSAPI) DeregisterTaskDefinition(_param0 *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error) {
	ret := _m.ctrl.Call(_m, "DeregisterTaskDefinition", _param0)
	ret0, _ := ret[0].(*ecs.DeregisterTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DeregisterTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeregisterTaskDefinition", arg0)
}

func (_m *MockECSAPI) DescribeClusters(_param0 *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeClusters", _param0)
	ret0, _ := ret[0].(*ecs.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeClusters(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeClusters", arg0)
}

func (_m *MockECSAPI) DescribeContainerInstances(_param0 *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeContainerInstances", _param0)
	ret0, _ := ret[0].(*ecs.DescribeContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeContainerInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeContainerInstances", arg0)
}

func (_m *MockECSAPI) DescribeServices(_param0 *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeServices", _param0)
	ret0, _ := ret[0].(*ecs.DescribeServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeServices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeServices", arg0)
}

func (_m *MockECSAPI) DescribeTaskDefinition(_param0 *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeTaskDefinition", _param0)
	ret0, _ := ret[0].(*ecs.DescribeTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTaskDefinition", arg0)
}

func (_m *MockECSAPI) DescribeTasks(_param0 *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeTasks", _param0)
	ret0, _ := ret[0].(*ecs.DescribeTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DescribeTasks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeTasks", arg0)
}

func (_m *MockECSAPI) DiscoverPollEndpoint(_param0 *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
	ret := _m.ctrl.Call(_m, "DiscoverPollEndpoint", _param0)
	ret0, _ := ret[0].(*ecs.DiscoverPollEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) DiscoverPollEndpoint(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiscoverPollEndpoint", arg0)
}

func (_m *MockECSAPI) ListClusters(_param0 *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	ret := _m.ctrl.Call(_m, "ListClusters", _param0)
	ret0, _ := ret[0].(*ecs.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListClusters(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListClusters", arg0)
}

func (_m *MockECSAPI) ListContainerInstances(_param0 *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListContainerInstances", _param0)
	ret0, _ := ret[0].(*ecs.ListContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListContainerInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainerInstances", arg0)
}

func (_m *MockECSAPI) ListServices(_param0 *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListServices", _param0)
	ret0, _ := ret[0].(*ecs.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListServices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListServices", arg0)
}

func (_m *MockECSAPI) ListTaskDefinitionFamilies(_param0 *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitionFamilies", _param0)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionFamiliesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitionFamilies(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitionFamilies", arg0)
}

func (_m *MockECSAPI) ListTaskDefinitions(_param0 *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	ret := _m.ctrl.Call(_m, "ListTaskDefinitions", _param0)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTaskDefinitions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTaskDefinitions", arg0)
}

func (_m *MockECSAPI) ListTasks(_param0 *ecs.ListTasksInput) (*ecs.ListTasksOutput, error) {
	ret := _m.ctrl.Call(_m, "ListTasks", _param0)
	ret0, _ := ret[0].(*ecs.ListTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) ListTasks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasks", arg0)
}

func (_m *MockECSAPI) RegisterContainerInstance(_param0 *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
	ret := _m.ctrl.Call(_m, "RegisterContainerInstance", _param0)
	ret0, _ := ret[0].(*ecs.RegisterContainerInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RegisterContainerInstance(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterContainerInstance", arg0)
}

func (_m *MockECSAPI) RegisterTaskDefinition(_param0 *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
	ret := _m.ctrl.Call(_m, "RegisterTaskDefinition", _param0)
	ret0, _ := ret[0].(*ecs.RegisterTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RegisterTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterTaskDefinition", arg0)
}

func (_m *MockECSAPI) RunTask(_param0 *ecs.RunTaskInput) (*ecs.RunTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "RunTask", _param0)
	ret0, _ := ret[0].(*ecs.RunTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) RunTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RunTask", arg0)
}

func (_m *MockECSAPI) StartTask(_param0 *ecs.StartTaskInput) (*ecs.StartTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "StartTask", _param0)
	ret0, _ := ret[0].(*ecs.StartTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) StartTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartTask", arg0)
}

func (_m *MockECSAPI) StopTask(_param0 *ecs.StopTaskInput) (*ecs.StopTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "StopTask", _param0)
	ret0, _ := ret[0].(*ecs.StopTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) StopTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopTask", arg0)
}

func (_m *MockECSAPI) SubmitContainerStateChange(_param0 *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
	ret := _m.ctrl.Call(_m, "SubmitContainerStateChange", _param0)
	ret0, _ := ret[0].(*ecs.SubmitContainerStateChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) SubmitContainerStateChange(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubmitContainerStateChange", arg0)
}

func (_m *MockECSAPI) SubmitTaskStateChange(_param0 *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
	ret := _m.ctrl.Call(_m, "SubmitTaskStateChange", _param0)
	ret0, _ := ret[0].(*ecs.SubmitTaskStateChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) SubmitTaskStateChange(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SubmitTaskStateChange", arg0)
}

func (_m *MockECSAPI) UpdateContainerAgent(_param0 *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateContainerAgent", _param0)
	ret0, _ := ret[0].(*ecs.UpdateContainerAgentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateContainerAgent(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateContainerAgent", arg0)
}

func (_m *MockECSAPI) UpdateService(_param0 *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	ret := _m.ctrl.Call(_m, "UpdateService", _param0)
	ret0, _ := ret[0].(*ecs.UpdateServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockECSAPIRecorder) UpdateService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateService", arg0)
}
