// Automatically generated by MockGen. DO NOT EDIT!
// Source: assert/exam.go

package assert

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Exam interface
type MockExam struct {
	ctrl     *gomock.Controller
	recorder *_MockExamRecorder
}

// Recorder for MockExam (not exported)
type _MockExamRecorder struct {
	mock *MockExam
}

func NewMockExam(ctrl *gomock.Controller) *MockExam {
	mock := &MockExam{ctrl: ctrl}
	mock.recorder = &_MockExamRecorder{mock}
	return mock
}

func (_m *MockExam) EXPECT() *_MockExamRecorder {
	return _m.recorder
}

func (_m *MockExam) Fail() {
	_m.ctrl.Call(_m, "Fail")
}

func (_mr *_MockExamRecorder) Fail() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fail")
}

func (_m *MockExam) FailNow() {
	_m.ctrl.Call(_m, "FailNow")
}

func (_mr *_MockExamRecorder) FailNow() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FailNow")
}

func (_m *MockExam) Log(args ...interface{}) {
	_s := []interface{}{}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Log", _s...)
}

func (_mr *_MockExamRecorder) Log(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Log", arg0...)
}

func (_m *MockExam) Logf(format string, args ...interface{}) {
	_s := []interface{}{format}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Logf", _s...)
}

func (_mr *_MockExamRecorder) Logf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Logf", _s...)
}
