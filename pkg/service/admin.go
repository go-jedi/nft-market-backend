package service

import (
	"github.com/rob-bender/nft-market-backend/appl_row"
	"github.com/rob-bender/nft-market-backend/pkg/repository"
)

type AdminService struct {
	repo repository.TodoAdmin
}

func NewAdminService(r repository.TodoAdmin) *AdminService {
	return &AdminService{
		repo: r,
	}
}

func (s *AdminService) CheckIsAdmin(teleId int64) (bool, int, error) {
	return s.repo.CheckIsAdmin(teleId)
}

func (s *AdminService) CreateReferral(referralForm appl_row.ReferralCreate) (int, error) {
	return s.repo.CreateReferral(referralForm)
}

func (s *AdminService) GetUsersReferral(teleId int64) ([]appl_row.Referral, int, error) {
	return s.repo.GetUsersReferral(teleId)
}

func (s *AdminService) AdminGetUserProfile(teleId int64) ([]appl_row.AdminUserProfileGet, int, error) {
	return s.repo.AdminGetUserProfile(teleId)
}

func (s *AdminService) UpdatePremium(teleId int64) (int, error) {
	return s.repo.UpdatePremium(teleId)
}
