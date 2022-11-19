package service

import (
	"github.com/rob-bender/nft-market-backend/appl_row"
	"github.com/rob-bender/nft-market-backend/pkg/repository"
)

type TodoUser interface {
	CheckAuth(teleId int64) (bool, int, error)
	RegistrationUser(userForm appl_row.UserCreate) (int, error)
	GetUserLanguage(teleId int64) ([]appl_row.UserLanguage, int, error)
	UpdateLanguage(userForm appl_row.UserUpdateLanguage) (int, error)
	GetUserCurrency(teleId int64) ([]appl_row.UserCurrency, int, error)
	UpdateCurrency(userFormCurrency appl_row.UserUpdateCurrency) (int, error)
	CheckIsTerms(teleId int64) (bool, int, error)
	AgreeTerms(teleId int64) (int, error)
	GetUserProfile(teleId int64) ([]appl_row.UserProfile, int, error)
}

type TodoPayment interface {
	CreatePayment(paymentForm appl_row.PaymentCreate) (int, error)
	GetAllPayments() ([]appl_row.Payment, int, error)
}

type TodoCollection interface {
	CreateCollection(collectionForm appl_row.CollectionCreate) (int, error)
	GetAllCollections() ([]appl_row.Collection, int, error)
}

type Service struct {
	TodoUser
	TodoPayment
	TodoCollection
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		TodoUser:       NewUserService(r.TodoUser),
		TodoPayment:    NewPaymentService(r.TodoPayment),
		TodoCollection: NewCollectionService(r.TodoCollection),
	}
}
