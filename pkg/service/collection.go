package service

import (
	"github.com/rob-bender/nft-market-backend/appl_row"
	"github.com/rob-bender/nft-market-backend/pkg/repository"
)

type CollectionService struct {
	repo repository.TodoCollection
}

func NewCollectionService(r repository.TodoCollection) *CollectionService {
	return &CollectionService{
		repo: r,
	}
}

func (s *CollectionService) CreateCollection(collectionForm appl_row.CollectionCreate) (int, error) {
	return s.repo.CreateCollection(collectionForm)
}

func (s *CollectionService) GetAllCollections() ([]appl_row.Collection, int, error) {
	return s.repo.GetAllCollections()
}
