package assert

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_IsEqInt_Unequal(t *testing.T) {
	finish, ex := NewExam(t)
	defer finish()

	a, b := 1, 2
	v := a == b
	ex.EXPECT().Logf(gomock.Any(), a, b, v).Times(1)
	ex.EXPECT().Fail().Times(1)

	IsEqInt(ex, 1, 2)
}

func Test_IsEqInt_Equal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ex := NewMockExam(ctrl)

	a, b := 2, 2
	IsEqInt(ex, a, b)
}

func Test_IsPos_False(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	a := -10
	ex := NewMockExam(ctrl)
	ex.EXPECT().Logf(gomock.Any(), a, 0, false).Times(1)
	ex.EXPECT().Fail().Times(1)

	IsPos(ex, a)
}

func Test_IsPos_Zero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	a := 0
	ex := NewMockExam(ctrl)
	ex.EXPECT().Logf(gomock.Any(), a, 0, false).Times(1)
	ex.EXPECT().Fail().Times(1)

	IsPos(ex, a)
}

func Test_IsPos_True(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ex := NewMockExam(ctrl)

	IsPos(ex, 10)
}

func Test_IsZero_NonZero(t *testing.T) {
	finish, ex := NewExam(t)
	defer finish()
	ex.EXPECT().Fail().Times(1)
	ex.EXPECT().Logf(gomock.Any(), 1, 0, 1 == 0).Times(1)

	IsZero(ex, 1)
}

func Test_IsZero_Zero(t *testing.T) {
	finish, ex := NewExam(t)
	defer finish()

	IsZero(ex, 0)
}

func NewExam(t *testing.T) (func(), *MockExam) {
	ctrl := gomock.NewController(t)
	ex := NewMockExam(ctrl)
	return ctrl.Finish, ex
}
