// DO NOT EDIT
// Code generated automatically by github.com/efritz/go-mockgen
// $ go-mockgen github.com/efritz/deepjoy -i Client -d mock

package mock

import (
	deepjoy "github.com/efritz/deepjoy"
	sync "sync"
)

type MockClient struct {
	statsCloseLock          sync.RWMutex
	statCloseFuncCallCount  int
	statCloseFuncCallParams []ClientCloseParamSet
	CloseFunc               func()

	statsDoLock          sync.RWMutex
	statDoFuncCallCount  int
	statDoFuncCallParams []ClientDoParamSet
	DoFunc               func(string, ...interface{}) (interface{}, error)

	statsReadReplicaLock          sync.RWMutex
	statReadReplicaFuncCallCount  int
	statReadReplicaFuncCallParams []ClientReadReplicaParamSet
	ReadReplicaFunc               func() deepjoy.Client

	statsTransactionLock          sync.RWMutex
	statTransactionFuncCallCount  int
	statTransactionFuncCallParams []ClientTransactionParamSet
	TransactionFunc               func(...deepjoy.Command) (interface{}, error)
}
type ClientCloseParamSet struct{}
type ClientDoParamSet struct {
	Arg0 string
	Arg1 []interface{}
}
type ClientReadReplicaParamSet struct{}
type ClientTransactionParamSet struct {
	Arg0 []deepjoy.Command
}

var _ deepjoy.Client = NewMockClient()

func NewMockClient() *MockClient {
	m := &MockClient{}
	m.DoFunc = m.defaultDoFunc
	m.ReadReplicaFunc = m.defaultReadReplicaFunc
	m.TransactionFunc = m.defaultTransactionFunc
	m.CloseFunc = m.defaultCloseFunc
	return m
}
func (m *MockClient) Close() {
	m.statsCloseLock.Lock()
	m.statCloseFuncCallCount++
	m.statCloseFuncCallParams = append(m.statCloseFuncCallParams, ClientCloseParamSet{})
	m.statsCloseLock.Unlock()
	m.CloseFunc()
}
func (m *MockClient) CloseFuncCallCount() int {
	m.statsCloseLock.RLock()
	defer m.statsCloseLock.RUnlock()
	return m.statCloseFuncCallCount
}
func (m *MockClient) CloseFuncCallParams() []ClientCloseParamSet {
	m.statsCloseLock.RLock()
	defer m.statsCloseLock.RUnlock()
	return m.statCloseFuncCallParams
}

func (m *MockClient) Do(v0 string, v1 ...interface{}) (interface{}, error) {
	m.statsDoLock.Lock()
	m.statDoFuncCallCount++
	m.statDoFuncCallParams = append(m.statDoFuncCallParams, ClientDoParamSet{v0, v1})
	m.statsDoLock.Unlock()
	return m.DoFunc(v0, v1...)
}
func (m *MockClient) DoFuncCallCount() int {
	m.statsDoLock.RLock()
	defer m.statsDoLock.RUnlock()
	return m.statDoFuncCallCount
}
func (m *MockClient) DoFuncCallParams() []ClientDoParamSet {
	m.statsDoLock.RLock()
	defer m.statsDoLock.RUnlock()
	return m.statDoFuncCallParams
}

func (m *MockClient) ReadReplica() deepjoy.Client {
	m.statsReadReplicaLock.Lock()
	m.statReadReplicaFuncCallCount++
	m.statReadReplicaFuncCallParams = append(m.statReadReplicaFuncCallParams, ClientReadReplicaParamSet{})
	m.statsReadReplicaLock.Unlock()
	return m.ReadReplicaFunc()
}
func (m *MockClient) ReadReplicaFuncCallCount() int {
	m.statsReadReplicaLock.RLock()
	defer m.statsReadReplicaLock.RUnlock()
	return m.statReadReplicaFuncCallCount
}
func (m *MockClient) ReadReplicaFuncCallParams() []ClientReadReplicaParamSet {
	m.statsReadReplicaLock.RLock()
	defer m.statsReadReplicaLock.RUnlock()
	return m.statReadReplicaFuncCallParams
}

func (m *MockClient) Transaction(v0 ...deepjoy.Command) (interface{}, error) {
	m.statsTransactionLock.Lock()
	m.statTransactionFuncCallCount++
	m.statTransactionFuncCallParams = append(m.statTransactionFuncCallParams, ClientTransactionParamSet{v0})
	m.statsTransactionLock.Unlock()
	return m.TransactionFunc(v0...)
}
func (m *MockClient) TransactionFuncCallCount() int {
	m.statsTransactionLock.RLock()
	defer m.statsTransactionLock.RUnlock()
	return m.statTransactionFuncCallCount
}
func (m *MockClient) TransactionFuncCallParams() []ClientTransactionParamSet {
	m.statsTransactionLock.RLock()
	defer m.statsTransactionLock.RUnlock()
	return m.statTransactionFuncCallParams
}

func (m *MockClient) defaultCloseFunc() {
	return
}
func (m *MockClient) defaultDoFunc(v0 string, v1 ...interface{}) (interface{}, error) {
	return nil, nil
}
func (m *MockClient) defaultReadReplicaFunc() deepjoy.Client {
	return nil
}
func (m *MockClient) defaultTransactionFunc(v0 ...deepjoy.Command) (interface{}, error) {
	return nil, nil
}
