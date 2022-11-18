package service

import (
	"github.com/rob-bender/our_project/appl_row"
	"github.com/rob-bender/our_project/pkg/repository"
)

type TodoUser interface {
	RegistrationUser(userForm appl_row.UserCreate) (error, int)
	CheckAuth(teleId int64) (bool, error, int)
	UpdateLanguage(userForm appl_row.UserUpdateLanguage) (error, int)
	CheckIsLanguage(teleId int64) ([]appl_row.CheckUserLanguageResponse, error, int)
	UpdateCurrency(userFormCurrency appl_row.UserUpdateCurrency) (error, int)
	CheckIsTerms(teleId int64) (bool, error, int)
	AgreeTerms(teleId int64) (error, int)
}

type Service struct {
	TodoUser
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoUser: NewUserService(r.TodoUser),
	}
}
