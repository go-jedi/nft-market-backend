package service

import (
	"github.com/rob-bender/nft-market-backend/appl_row"
	"github.com/rob-bender/nft-market-backend/pkg/repository"
)

type DepotService struct {
	repo repository.TodoDepot
}

func NewDepotService(r repository.TodoDepot) *DepotService {
	return &DepotService{
		repo: r,
	}
}

func (s *DepotService) CreateDepot(depotForm appl_row.DepotCreate) (int, error) {
	return s.repo.CreateDepot(depotForm)
}
