package service

import (
	"fmt"

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

// Create checks that list belongs to user and creates new item.
// It returns errors if something wrong happened.
func (svc *ItemService) Create(userID, listID int64, opts CreateItemOpts) error {
	_, err := svc.ListRepo.GetByID(userID, listID)
	if err != nil {
		return fmt.Errorf("ItemService: %v", err)
	}

	err = svc.ItemRepo.Create(listID, repository.CreateItemOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})

	if err != nil {
		return fmt.Errorf("ItemService: %v", err)
	}

	return nil
}

// GetAll returns all items of user list or
// Errors is something wrong happened.
func (svc *ItemService) GetAll(userID, listID int64) ([]domain.Item, error) {
	items, err := svc.ItemRepo.GetAll(userID, listID)
	if err != nil {
		return nil, fmt.Errorf("ItemService: %v", err)
	}

	result := make([]domain.Item, len(items))
	for i, item := range items {
		result[i] = *item.ToDomain()
	}

	return result, nil
}

// GetByID returns user item with passed id or
// Errors if something wrong happened.
func (svc *ItemService) GetByID(userID, itemID int64) (*domain.Item, error) {
	item, err := svc.ItemRepo.GetByID(userID, itemID)
	if err != nil {
		return nil, fmt.Errorf("ItemService: %v", err)
	}

	return item.ToDomain(), nil
}

type UpdateItemOpts struct {
	Title       *string
	Description *string
	Done        *bool
}

// Update updates user item with passed id
// It returns errors if something wrong happened.
func (svc *ItemService) Update(userID, itemID int64, opts UpdateItemOpts) error {
	err := svc.ItemRepo.Update(userID, itemID, repository.UpdateItemOpts{
		Title:       opts.Title,
		Description: opts.Description,
		Done:        opts.Done,
	})
	if err != nil {
		return fmt.Errorf("ItemService: %v", err)
	}

	return nil
}

// DeleteByID removes user item with passed id
// It returns errors if something wrong happened.
func (svc *ItemService) DeleteByID(userID, itemID int64) error {
	err := svc.ItemRepo.DeleteByID(userID, itemID)
	if err != nil {
		return fmt.Errorf("IteService: %v", err)
	}

	return nil
}
