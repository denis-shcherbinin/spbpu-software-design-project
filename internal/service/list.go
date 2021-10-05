package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
)

type ListService struct {
	ListRepo repository.List
}

func NewListService(listRepo repository.List) *ListService {
	return &ListService{
		ListRepo: listRepo,
	}
}

type CreateListOpts struct {
	UserID      int64
	Title       string
	Description string
}

func (svc *ListService) Create(opts CreateListOpts) error {
	return svc.ListRepo.Create(repository.CreateListOpts{
		UserID:      opts.UserID,
		Title:       opts.Title,
		Description: opts.Description,
	})
}

func (svc *ListService) GetAll(userID int64) ([]domain.List, error) {
	lists, err := svc.ListRepo.GetAll(userID)
	if err != nil {
		return nil, err
	}

	result := make([]domain.List, len(lists))
	for i, l := range lists {
		result[i] = *l.ToDomain()
	}

	return result, nil
}

func (svc *ListService) GetByID(userID, listID int64) (*domain.List, error) {
	list, err := svc.ListRepo.GetByID(userID, listID)
	if err != nil {
		return nil, err
	}

	return list.ToDomain(), nil
}

type UpdateListOpts struct {
	Title       string
	Description string
}

func (svc *ListService) Update(userID, listID int64, opts UpdateListOpts) error {
	return svc.ListRepo.Update(userID, listID, repository.UpdateListOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})
}

func (svc *ListService) DeleteByID(userID, listID int64) error {
	return svc.ListRepo.DeleteByID(userID, listID)
}
