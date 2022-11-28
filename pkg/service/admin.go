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

func (s *AdminService) CheckUserReferral(teleId int64) ([]appl_row.CheckUserReferralGet, int, error) {
	return s.repo.CheckUserReferral(teleId)
}

func (s *AdminService) GetUserReferral(teleId int64, teleIdUser int64) ([]appl_row.Referral, int, error) {
	return s.repo.GetUserReferral(teleId, teleIdUser)
}

func (s *AdminService) GetUsersReferral(teleId int64, limit int) ([]appl_row.Referral, int, error) {
	return s.repo.GetUsersReferral(teleId, limit)
}

func (s *AdminService) AdminGetUserProfile(teleId int64) ([]appl_row.AdminUserProfileGet, int, error) {
	return s.repo.AdminGetUserProfile(teleId)
}

func (s *AdminService) CheckIsPremium(teleId int64) (bool, int, error) {
	return s.repo.CheckIsPremium(teleId)
}

func (s *AdminService) UpdatePremium(teleId int64) (int, error) {
	return s.repo.UpdatePremium(teleId)
}

func (s *AdminService) CheckIsVerification(teleId int64) (bool, int, error) {
	return s.repo.CheckIsVerification(teleId)
}

func (s *AdminService) UpdateVerification(teleId int64) (int, error) {
	return s.repo.UpdateVerification(teleId)
}

func (s *AdminService) AdminUpdateMinimPrice(teleId int64, minPrice float64) (int, error) {
	return s.repo.AdminUpdateMinimPrice(teleId, minPrice)
}

func (s *AdminService) AdminAddBalance(teleId int64, needPrice float64) (int, error) {
	return s.repo.AdminAddBalance(teleId, needPrice)
}

func (s *AdminService) AdminChangeMinUser(teleId int64, minPrice float64) (int, error) {
	return s.repo.AdminChangeMinUser(teleId, minPrice)
}

func (s *AdminService) AdminChangeBalance(teleId int64, needPrice float64) (int, error) {
	return s.repo.AdminChangeBalance(teleId, needPrice)
}

func (s *AdminService) CheckIsBlockUser(teleId int64) (bool, int, error) {
	return s.repo.CheckIsBlockUser(teleId)
}

func (s *AdminService) AdminBlockUser(teleId int64) (int, error) {
	return s.repo.AdminBlockUser(teleId)
}

func (s *AdminService) CheckIsVisibleName(teleId int64) (bool, int, error) {
	return s.repo.CheckIsVisibleName(teleId)
}

func (s *AdminService) AdminChangeVisibleName(teleId int64) (int, error) {
	return s.repo.AdminChangeVisibleName(teleId)
}

func (s *AdminService) AdminBuyTokenUser(teleId int64, tokenUid string, priceUser float64, uidPaymentEvent string) (int, error) {
	return s.repo.AdminBuyTokenUser(teleId, tokenUid, priceUser, uidPaymentEvent)
}

func (s *AdminService) AdminWithdrawApprove(teleId int64, withDrawEventUid string) (bool, int, error) {
	return s.repo.AdminWithdrawApprove(teleId, withDrawEventUid)
}
