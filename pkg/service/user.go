package service

import (
	"github.com/rob-bender/nft-market-backend/appl_row"
	"github.com/rob-bender/nft-market-backend/pkg/repository"
)

type UserService struct {
	repo repository.TodoUser
}

func NewUserService(r repository.TodoUser) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) GetAllExchangeRates() ([]appl_row.ExchangeRatesGet, int, error) {
	return s.repo.GetAllExchangeRates()
}

func (s *UserService) CheckAuth(teleId int64) (bool, int, error) {
	return s.repo.CheckAuth(teleId)
}

func (s *UserService) RegistrationUser(userForm appl_row.UserCreate) (int, error) {
	return s.repo.RegistrationUser(userForm)
}

func (s *UserService) GetUserLanguage(teleId int64) ([]appl_row.UserLanguage, int, error) {
	return s.repo.GetUserLanguage(teleId)
}

func (s *UserService) UpdateLanguage(userForm appl_row.UserUpdateLanguage) (int, error) {
	return s.repo.UpdateLanguage(userForm)
}

func (s *UserService) GetUserCurrency(teleId int64) ([]appl_row.UserCurrency, int, error) {
	return s.repo.GetUserCurrency(teleId)
}

func (s *UserService) UpdateCurrency(userFormCurrency appl_row.UserUpdateCurrency) (int, error) {
	return s.repo.UpdateCurrency(userFormCurrency)
}

func (s *UserService) CheckIsTerms(teleId int64) (bool, int, error) {
	return s.repo.CheckIsTerms(teleId)
}

func (s *UserService) AgreeTerms(teleId int64) (int, error) {
	return s.repo.AgreeTerms(teleId)
}

func (s *UserService) GetUserProfile(teleId int64) ([]appl_row.UserProfile, int, error) {
	return s.repo.GetUserProfile(teleId)
}

func (s *UserService) GetUserMinPrice(teleId int64) ([]appl_row.UserMinPrice, int, error) {
	return s.repo.GetUserMinPrice(teleId)
}

func (s *UserService) GetAdminByUser(teleId int64) ([]appl_row.AdminByUser, int, error) {
	return s.repo.GetAdminByUser(teleId)
}

func (s *UserService) GetUserBalance(teleId int64) ([]appl_row.UserBalance, int, error) {
	return s.repo.GetUserBalance(teleId)
}

func (s *UserService) CheckUserToken(teleId int64, tokenUid string) (bool, int, error) {
	return s.repo.CheckUserToken(teleId, tokenUid)
}

func (s *UserService) BuyUserToken(userBuyTokenForm appl_row.UserBuyToken) (int, error) {
	return s.repo.BuyUserToken(userBuyTokenForm)
}

func (s *UserService) SellUserToken(userSellTokenForm appl_row.UserSellToken) (string, int, error) {
	return s.repo.SellUserToken(userSellTokenForm)
}

func (s *UserService) GetUserNft(teleId int64) (appl_row.UserGetNft, int, error) {
	return s.repo.GetUserNft(teleId)
}

func (s *UserService) GetUserPaymentEvent(eventUid string) ([]appl_row.UserGetUserPaymentEvent, int, error) {
	return s.repo.GetUserPaymentEvent(eventUid)
}

func (s *UserService) CreateWithDrawEvent(userWithDrawForm appl_row.UserWithDrawEventCreate) (string, int, error) {
	return s.repo.CreateWithDrawEvent(userWithDrawForm)
}

func (s *UserService) GetWithDrawEvent(withDrawEventUid string) ([]appl_row.WithDrawEventGet, int, error) {
	return s.repo.GetWithDrawEvent(withDrawEventUid)
}
