package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
)

type ItemService struct {
	ItemRepo repository.Item
	ListRepo repository.List
}

func NewItemService(itemRepo repository.Item, listRepo repository.List) *ItemService {
	return &ItemService{
		ItemRepo: itemRepo,
		ListRepo: listRepo,
	}
}

type CreateItemOpts struct {
	Title       string
	Description string
}

// Create .
func (svc *ItemService) Create(userID, listID int64, opts CreateItemOpts) error {
	_, err := svc.ListRepo.GetByID(userID, listID)
	if err != nil {
		return err
	}

	err = svc.ItemRepo.Create(listID, repository.CreateItemOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

// GetAll .
func (svc *ItemService) GetAll(userID, listID int64) ([]domain.Item, error) {
	items, err := svc.ItemRepo.GetAll(userID, listID)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Item, len(items))
	for i, item := range items {
		result[i] = *item.ToDomain()
	}

	return result, nil
}

// GetByID .
func (svc *ItemService) GetByID(userID, itemID int64) (*domain.Item, error) {
	item, err := svc.ItemRepo.GetByID(userID, itemID)
	if err != nil {
		return nil, err
	}

	return item.ToDomain(), nil
}

type UpdateItemOpts struct {
	Title       *string
	Description *string
	Done        *bool
}

// Update .
func (svc *ItemService) Update(userID, itemID int64, opts UpdateItemOpts) error {
	err := svc.ItemRepo.Update(userID, itemID, repository.UpdateItemOpts{
		Title:       opts.Title,
		Description: opts.Description,
		Done:        opts.Done,
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID .
func (svc *ItemService) DeleteByID(userID, itemID int64) error {
	err := svc.ItemRepo.DeleteByID(userID, itemID)
	if err != nil {
		return err
	}

	return nil
}
