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

func (s *CollectionService) CreateToken(tokenForm appl_row.TokenCreate) (int, error) {
	return s.repo.CreateToken(tokenForm)
}

func (s *CollectionService) GetAllTokensCollection(uidCollection string) ([]appl_row.TokensGetByCollection, int, error) {
	return s.repo.GetAllTokensCollection(uidCollection)
}

func (s *CollectionService) GetToken(tokenUid string) ([]appl_row.Token, int, error) {
	return s.repo.GetToken(tokenUid)
}
