// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/run/cmd.go

// Package run is a generated GoMock package.
package run

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	os "os"
	reflect "reflect"
	syscall "syscall"
)

// Mockcmd is a mock of cmd interface
type Mockcmd struct {
	ctrl     *gomock.Controller
	recorder *MockcmdMockRecorder
}

// MockcmdMockRecorder is the mock recorder for Mockcmd
type MockcmdMockRecorder struct {
	mock *Mockcmd
}

// NewMockcmd creates a new mock instance
func NewMockcmd(ctrl *gomock.Controller) *Mockcmd {
	mock := &Mockcmd{ctrl: ctrl}
	mock.recorder = &MockcmdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockcmd) EXPECT() *MockcmdMockRecorder {
	return m.recorder
}

// Output mocks base method
func (m *Mockcmd) Output() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Output")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Output indicates an expected call of Output
func (mr *MockcmdMockRecorder) Output() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*Mockcmd)(nil).Output))
}

// Run mocks base method
func (m *Mockcmd) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockcmdMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*Mockcmd)(nil).Run))
}

// CombinedOutput mocks base method
func (m *Mockcmd) CombinedOutput() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CombinedOutput")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CombinedOutput indicates an expected call of CombinedOutput
func (mr *MockcmdMockRecorder) CombinedOutput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CombinedOutput", reflect.TypeOf((*Mockcmd)(nil).CombinedOutput))
}

// Start mocks base method
func (m *Mockcmd) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockcmdMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockcmd)(nil).Start))
}

// StderrPipe mocks base method
func (m *Mockcmd) StderrPipe() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StderrPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StderrPipe indicates an expected call of StderrPipe
func (mr *MockcmdMockRecorder) StderrPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StderrPipe", reflect.TypeOf((*Mockcmd)(nil).StderrPipe))
}

// StdinPipe mocks base method
func (m *Mockcmd) StdinPipe() (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdinPipe")
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StdinPipe indicates an expected call of StdinPipe
func (mr *MockcmdMockRecorder) StdinPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdinPipe", reflect.TypeOf((*Mockcmd)(nil).StdinPipe))
}

// StdoutPipe mocks base method
func (m *Mockcmd) StdoutPipe() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdoutPipe")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StdoutPipe indicates an expected call of StdoutPipe
func (mr *MockcmdMockRecorder) StdoutPipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdoutPipe", reflect.TypeOf((*Mockcmd)(nil).StdoutPipe))
}

// String mocks base method
func (m *Mockcmd) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockcmdMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*Mockcmd)(nil).String))
}

// Wait mocks base method
func (m *Mockcmd) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockcmdMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*Mockcmd)(nil).Wait))
}

// Path mocks base method
func (m *Mockcmd) Path(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Path", arg0)
}

// Path indicates an expected call of Path
func (mr *MockcmdMockRecorder) Path(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*Mockcmd)(nil).Path), arg0)
}

// Args mocks base method
func (m *Mockcmd) Args(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Args", arg0)
}

// Args indicates an expected call of Args
func (mr *MockcmdMockRecorder) Args(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Args", reflect.TypeOf((*Mockcmd)(nil).Args), arg0)
}

// Env mocks base method
func (m *Mockcmd) Env(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Env", arg0)
}

// Env indicates an expected call of Env
func (mr *MockcmdMockRecorder) Env(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Env", reflect.TypeOf((*Mockcmd)(nil).Env), arg0)
}

// Dir mocks base method
func (m *Mockcmd) Dir(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dir", arg0)
}

// Dir indicates an expected call of Dir
func (mr *MockcmdMockRecorder) Dir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dir", reflect.TypeOf((*Mockcmd)(nil).Dir), arg0)
}

// Stdin mocks base method
func (m *Mockcmd) Stdin(arg0 io.Reader) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stdin", arg0)
}

// Stdin indicates an expected call of Stdin
func (mr *MockcmdMockRecorder) Stdin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stdin", reflect.TypeOf((*Mockcmd)(nil).Stdin), arg0)
}

// Stdout mocks base method
func (m *Mockcmd) Stdout(arg0 io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stdout", arg0)
}

// Stdout indicates an expected call of Stdout
func (mr *MockcmdMockRecorder) Stdout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stdout", reflect.TypeOf((*Mockcmd)(nil).Stdout), arg0)
}

// Stderr mocks base method
func (m *Mockcmd) Stderr(arg0 io.Writer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stderr", arg0)
}

// Stderr indicates an expected call of Stderr
func (mr *MockcmdMockRecorder) Stderr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stderr", reflect.TypeOf((*Mockcmd)(nil).Stderr), arg0)
}

// ExtraFiles mocks base method
func (m *Mockcmd) ExtraFiles(arg0 []*os.File) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExtraFiles", arg0)
}

// ExtraFiles indicates an expected call of ExtraFiles
func (mr *MockcmdMockRecorder) ExtraFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtraFiles", reflect.TypeOf((*Mockcmd)(nil).ExtraFiles), arg0)
}

// SysProcAttr mocks base method
func (m *Mockcmd) SysProcAttr(arg0 *syscall.SysProcAttr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SysProcAttr", arg0)
}

// SysProcAttr indicates an expected call of SysProcAttr
func (mr *MockcmdMockRecorder) SysProcAttr(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SysProcAttr", reflect.TypeOf((*Mockcmd)(nil).SysProcAttr), arg0)
}

// Process mocks base method
func (m *Mockcmd) Process() *os.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process")
	ret0, _ := ret[0].(*os.Process)
	return ret0
}

// Process indicates an expected call of Process
func (mr *MockcmdMockRecorder) Process() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*Mockcmd)(nil).Process))
}

// ProcessState mocks base method
func (m *Mockcmd) ProcessState() *os.ProcessState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessState")
	ret0, _ := ret[0].(*os.ProcessState)
	return ret0
}

// ProcessState indicates an expected call of ProcessState
func (mr *MockcmdMockRecorder) ProcessState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessState", reflect.TypeOf((*Mockcmd)(nil).ProcessState))
}