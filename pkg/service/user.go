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
