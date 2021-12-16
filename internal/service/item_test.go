package service

import (
	"errors"
	"runtime/debug"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository/entity"
)

func TestItemService_Create(t *testing.T) {
	t.Parallel()

	type testEnv struct {
		itemSvc      Item
		itemMockRepo *repository.MockItem
		listMockRepo *repository.MockList
	}

	var newTestEnv = func(t *testing.T) *testEnv {
		var te = &testEnv{}
		var ctrl = gomock.NewController(t)

		te.itemMockRepo = repository.NewMockItem(ctrl)
		te.listMockRepo = repository.NewMockList(ctrl)

		te.itemSvc = NewItemService(te.itemMockRepo, te.listMockRepo)

		return te
	}

	type testCase struct {
		name   string
		setUp  func(te *testEnv) (userID, listID int64, opts CreateItemOpts)
		expect func() (err error)
		trace  func(t *testing.T)
	}

	var tests = []testCase{
		{
			name:  "OK",
			trace: func(t *testing.T) { t.Log(string(debug.Stack())) },
			setUp: func(te *testEnv) (userID, listID int64, opts CreateItemOpts) {
				userID = 1
				listID = 1

				opts = CreateItemOpts{
					Title:       "title",
					Description: "description",
				}

				te.listMockRepo.EXPECT().GetByID(userID, listID).
					Return(&entity.List{}, nil)

				te.itemMockRepo.EXPECT().Create(listID, repository.CreateItemOpts{
					Title:       opts.Title,
					Description: opts.Description,
				}).
					Return(nil)

				return userID, listID, opts
			},
			expect: func() (err error) {
				return nil
			},
		},
		{
			name: "ListNotFound error",
			trace: func(t *testing.T) {t.Log(string(debug.Stack()))},
			setUp: func(te *testEnv) (userID, listID int64, opts CreateItemOpts) {
				userID = 0
				listID = 0

				opts = CreateItemOpts{
					Title:       "title",
					Description: "description",
				}

				te.listMockRepo.EXPECT().GetByID(userID, listID).
					Return(nil, errs.ErrListNotFound)

				return int64(0), int64(0), opts
			},
			expect: func() (err error) {
				return errs.ErrListNotFound
			},
		},
	}

	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			var test = el
			t.Parallel()

			var te = newTestEnv(t)

			var s1, s2, s3 = test.setUp(te)

			var errExp = test.expect()

			var r1 = te.itemSvc.Create(s1, s2, s3)

			if !errors.Is(r1, errExp) {
				test.trace(t)
				t.Errorf("\nexp: %v\nget:%v", errExp, r1)
			}
		})
	}
}
