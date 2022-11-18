package service

import (
	"github.com/rob-bender/our_project/appl_row"
	"github.com/rob-bender/our_project/pkg/repository"
)

type UserService struct {
	repo repository.TodoUser
}

func NewUserService(r repository.TodoUser) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) RegistrationUser(userForm appl_row.UserCreate) (error, int) {
	return s.repo.RegistrationUser(userForm)
}

func (s *UserService) CheckAuth(teleId int64) (bool, error, int) {
	return s.repo.CheckAuth(teleId)
}

func (s *UserService) UpdateLanguage(userForm appl_row.UserUpdateLanguage) (error, int) {
	return s.repo.UpdateLanguage(userForm)
}

func (s *UserService) CheckIsLanguage(teleId int64) ([]appl_row.CheckUserLanguageResponse, error, int) {
	return s.repo.CheckIsLanguage(teleId)
}

func (s *UserService) UpdateCurrency(userFormCurrency appl_row.UserUpdateCurrency) (error, int) {
	return s.repo.UpdateCurrency(userFormCurrency)
}

func (s *UserService) CheckIsTerms(teleId int64) (bool, error, int) {
	return s.repo.CheckIsTerms(teleId)
}

func (s *UserService) AgreeTerms(teleId int64) (error, int) {
	return s.repo.AgreeTerms(teleId)
}
