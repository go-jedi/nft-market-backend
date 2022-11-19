package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nft-market-backend/appl_row"
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

type Repository struct {
	TodoUser
	TodoPayment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser:    NewUserPostgres(db),
		TodoPayment: NewPaymentPostgres(db),
	}
}
