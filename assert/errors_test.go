package assert

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestError_Nil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ex := NewMockExam(ctrl)

	IsNil(ex, nil)
}

func TestError_ProvidedNonNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ex := NewMockExam(ctrl)
	ex.EXPECT().Logf(gomock.Any(), ErrForTesting, ErrForTesting == nil).Times(1)
	ex.EXPECT().Fail().Times(1)

	IsNil(ex, ErrForTesting)
}

func Test_IsNil_StringNil(t *testing.T) {
	if isNil("<nil>") {
		t.Log("The string '<nil>' should not be considered nil")
		t.Fail()
	}
}

func TestError_NilInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Log("TestErrorNilInterface")
	ex := NewMockExam(ctrl)
	var f interface{}

	IsNil(ex, f)
}
