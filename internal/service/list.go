package service

import "github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"

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
