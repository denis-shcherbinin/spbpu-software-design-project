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
	Title       string
	Description string
}

// Create .
func (svc *ListService) Create(userID int64, opts CreateListOpts) error {
	err := svc.ListRepo.Create(userID, repository.CreateListOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

// GetAll .
func (svc *ListService) GetAll(userID int64) ([]domain.List, error) {
	lists, err := svc.ListRepo.GetAll(userID)
	if err != nil {
		return nil, err
	}

	result := make([]domain.List, len(lists))
	for i, list := range lists {
		result[i] = *list.ToDomain()
	}

	return result, nil
}

// GetByID .
func (svc *ListService) GetByID(userID, listID int64) (*domain.List, error) {
	list, err := svc.ListRepo.GetByID(userID, listID)
	if err != nil {
		return nil, err
	}

	return list.ToDomain(), nil
}

type UpdateListOpts struct {
	Title       *string
	Description *string
}

// Update .
func (svc *ListService) Update(userID, listID int64, opts UpdateListOpts) error {
	err := svc.ListRepo.Update(userID, listID, repository.UpdateListOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID .
func (svc *ListService) DeleteByID(userID, listID int64) error {
	err := svc.ListRepo.DeleteByID(userID, listID)
	if err != nil {
		return err
	}

	return nil
}
