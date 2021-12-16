package service

import (
	"errors"
	"runtime/debug"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
)

func TestUserService_GetIDByCredentials(t *testing.T) {
	t.Parallel()

	type testEnv struct {
		userSvc      User
		userRepoMock *repository.MockUser
	}

	var newTestEnv = func(t *testing.T) *testEnv {
		var te = &testEnv{}
		var ctrl = gomock.NewController(t)

		te.userRepoMock = repository.NewMockUser(ctrl)

		te.userSvc = NewUserService(te.userRepoMock)

		return te
	}

	type testCase struct {
		name   string
		setUp  func(te *testEnv) (username, passwordHash string)
		expect func() (id int64, err error)
		trace  func(t *testing.T)
	}

	var tests = []testCase{
		{
			name: "OK",
			trace: func(t *testing.T) { t.Log(string(debug.Stack()))},
			setUp: func(te *testEnv) (username, passwordHash string) {
				te.userRepoMock.EXPECT().GetIDByCredentials(username, passwordHash).
					Return(int64(1), nil)

				return username, passwordHash
			},
			expect: func() (id int64, err error) {
				return int64(1), nil
			},
		},
		{
			name: "UserNotFound error",
			trace: func(t *testing.T) {t.Log(string(debug.Stack()))},
			setUp: func(te *testEnv) (username, passwordHash string) {
				username = "username"
				passwordHash = "password_hash"

				te.userRepoMock.EXPECT().GetIDByCredentials(username, passwordHash).
					Return(int64(0), errs.ErrUserNotFound)

				return username, passwordHash
			},
			expect: func() (id int64, err error) {
				return int64(0), errs.ErrUserNotFound
			},
		},
	}

	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			var test = el
			t.Parallel()

			var te = newTestEnv(t)

			var s1,s2 = test.setUp(te)

			// It's not necessary to check id
			var _, errExp = test.expect()

			var _, r2 = te.userSvc.GetIDByCredentials(s1, s2)

			if !errors.Is(r2, errExp) {
				test.trace(t)
				t.Errorf("\nexp: %v\nget:%v", errExp, r2)
			}
		})
	}
}
