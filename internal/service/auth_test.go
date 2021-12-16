package service

import (
	"errors"
	"runtime/debug"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

const passwordSalt = "test_pwd_salt"

func TestAuthService_SignUp(t *testing.T) {
	t.Parallel()

	type testEnv struct {
		authSvc      Auth
		authRepoMock *repository.MockAuth
		userRepoMock *repository.MockUser
		hasher       hasher.Hasher
	}

	var newTestEnv = func(t *testing.T) *testEnv {
		var te = &testEnv{}
		var ctrl = gomock.NewController(t)

		te.authRepoMock = repository.NewMockAuth(ctrl)
		te.userRepoMock = repository.NewMockUser(ctrl)
		te.hasher = hasher.NewSHA1Hasher(passwordSalt)

		te.authSvc = NewAuthService(NewAuthOpts{
			AuthRepo: te.authRepoMock,
			UserRepo: te.userRepoMock,
			Hasher:   te.hasher,
		})

		return te
	}

	type testCase struct {
		name   string
		setUp  func(te *testEnv) (signUpOpts SignUpOpts)
		expect func() (err error)
		trace  func(t *testing.T)
	}

	var tests = []testCase{
		{
			name:  "OK",
			trace: func(t *testing.T) { t.Log(string(debug.Stack())) },
			setUp: func(te *testEnv) (signUpOpts SignUpOpts) {
				signUpOpts = SignUpOpts{
					FirstName:  "first_name",
					SecondName: "second_name",
					Username:   "username",
					Password:   "password",
				}

				te.authRepoMock.EXPECT().CreateUser(repository.CreateUserOpts{
					FirstName:  signUpOpts.FirstName,
					SecondName: signUpOpts.SecondName,
					Username:   signUpOpts.Username,
					Password:   te.hasher.Hash(signUpOpts.Password),
				}).
					Return(nil)

				return signUpOpts
			},
			expect: func() (err error) {
				return nil
			},
		},
		{
			name:  "UserAlreadyExists error",
			trace: func(t *testing.T) { t.Log(string(debug.Stack())) },
			setUp: func(te *testEnv) (signUpOpts SignUpOpts) {
				signUpOpts = SignUpOpts{
					FirstName:  "first_name",
					SecondName: "second_name",
					Username:   "username",
					Password:   "password",
				}

				te.authRepoMock.EXPECT().CreateUser(repository.CreateUserOpts{
					FirstName:  signUpOpts.FirstName,
					SecondName: signUpOpts.SecondName,
					Username:   signUpOpts.Username,
					Password:   te.hasher.Hash(signUpOpts.Password),
				}).
					Return(errs.ErrUserAlreadyExists)

				return signUpOpts
			},
			expect: func() (err error) {
				return errs.ErrUserAlreadyExists
			},
		},
	}

	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			var test = el
			t.Parallel()

			var te = newTestEnv(t)

			var s1 = test.setUp(te)

			var errExp = test.expect()

			var r1 = te.authSvc.SignUp(s1)

			if !errors.Is(r1, errExp) {
				test.trace(t)
				t.Errorf("\nexp: %v\nget:%v", errExp, r1)
			}
		})
	}
}

func TestAuthService_SignIn(t *testing.T) {
	t.Parallel()

	type testEnv struct {
		authSvc      Auth
		authRepoMock *repository.MockAuth
		userRepoMock *repository.MockUser
		hasher       hasher.Hasher
	}

	var newTestEnv = func(t *testing.T) *testEnv {
		var te = &testEnv{}
		var ctrl = gomock.NewController(t)

		te.authRepoMock = repository.NewMockAuth(ctrl)
		te.userRepoMock = repository.NewMockUser(ctrl)
		te.hasher = hasher.NewSHA1Hasher(passwordSalt)

		te.authSvc = NewAuthService(NewAuthOpts{
			AuthRepo: te.authRepoMock,
			UserRepo: te.userRepoMock,
			Hasher:   te.hasher,
		})

		return te
	}

	type testCase struct {
		name   string
		setUp  func(te *testEnv) (signInOpts SignInOpts)
		expect func() (username string, passwordHash string, err error)
		trace  func(t *testing.T)
	}

	var tests = []testCase{
		{
			name:  "OK",
			trace: func(t *testing.T) { t.Log(string(debug.Stack())) },
			setUp: func(te *testEnv) (signInOpts SignInOpts) {
				signInOpts = SignInOpts{
					Username: "username",
					Password: "password",
				}

				var passwordHash = te.hasher.Hash(signInOpts.Password)

				te.userRepoMock.EXPECT().CheckByCredentials(signInOpts.Username, passwordHash).
					Return(true, nil)

				return signInOpts
			},
			expect: func() (username string, passwordHash string, err error) {
				return "username", "password_hash", nil
			},
		},
		{
			name:  "UserNotFound error",
			trace: func(t *testing.T) { t.Log(string(debug.Stack())) },
			setUp: func(te *testEnv) (signInOpts SignInOpts) {
				signInOpts = SignInOpts{
					Username: "username",
					Password: "password",
				}

				var passwordHash = te.hasher.Hash(signInOpts.Password)

				te.userRepoMock.EXPECT().CheckByCredentials(signInOpts.Username, passwordHash).
					Return(false, nil)

				return signInOpts
			},
			expect: func() (username string, passwordHash string, err error) {
				return "", "", errs.ErrUserNotFound
			},
		},
	}

	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			var test = el
			t.Parallel()

			var te = newTestEnv(t)

			var s1 = test.setUp(te)

			// It's not necessary to check username or password hash
			var _, _, errExp = test.expect()

			var _, _, r3 = te.authSvc.SignIn(s1)

			if !errors.Is(r3, errExp) {
				test.trace(t)
				t.Errorf("\nexp: %v\nget:%v", errExp, r3)
			}
		})
	}
}
