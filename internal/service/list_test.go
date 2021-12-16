package service

import (
	"errors"
	"runtime/debug"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
)

func TestListService_Create(t *testing.T) {
	t.Parallel()

	type testEnv struct {
		listSvc      List
		listMockRepo *repository.MockList
	}

	var newTestEnv = func(t *testing.T) *testEnv {
		var te = &testEnv{}
		var ctrl = gomock.NewController(t)

		te.listMockRepo = repository.NewMockList(ctrl)

		te.listSvc = NewListService(te.listMockRepo)

		return te
	}

	type testCase struct {
		name   string
		setUp  func(te *testEnv) (userID int64, opts CreateListOpts)
		expect func() (err error)
		trace  func(t *testing.T)
	}

	var tests = []testCase{
		{
			name:  "OK",
			trace: func(t *testing.T) { t.Log(string(debug.Stack())) },
			setUp: func(te *testEnv) (userID int64, opts CreateListOpts) {
				userID = 1

				opts = CreateListOpts{
					Title:       "title",
					Description: "description",
				}

				te.listMockRepo.EXPECT().Create(userID, repository.CreateListOpts{
					Title:       opts.Title,
					Description: opts.Description,
				}).
					Return(nil)

				return userID, opts
			},
			expect: func() (err error) {
				return nil
			},
		},
	}

	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			var test = el
			t.Parallel()

			var te = newTestEnv(t)

			var s1, s2 = test.setUp(te)

			var errExp = test.expect()

			var r1 = te.listSvc.Create(s1, s2)

			if !errors.Is(r1, errExp) {
				test.trace(t)
				t.Errorf("\nexp: %v\nget:%v", errExp, r1)
			}
		})
	}
}
