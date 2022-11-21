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

type TodoCollection interface {
	CreateCollection(collectionForm appl_row.CollectionCreate) (int, error)
	GetAllCollections() ([]appl_row.Collection, int, error)
	CreateToken(tokenForm appl_row.TokenCreate) (int, error)
	GetAllTokensCollection(uidCollection string) ([]appl_row.TokensGetByCollection, int, error)
	GetToken(tokenUid string) ([]appl_row.Token, int, error)
}

type TodoAdmin interface {
	CheckIsAdmin(teleId int64) (bool, int, error)
	CreateReferral(referralForm appl_row.ReferralCreate) (int, error)
	GetUsersReferral(teleId int64) ([]appl_row.Referral, int, error)
	AdminGetUserProfile(teleId int64) ([]appl_row.AdminUserProfileGet, int, error)
	UpdatePremium(teleId int64) (int, error)
}

type Repository struct {
	TodoUser
	TodoPayment
	TodoCollection
	TodoAdmin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser:       NewUserPostgres(db),
		TodoPayment:    NewPaymentPostgres(db),
		TodoCollection: NewCollectionPostgres(db),
		TodoAdmin:      NewAdminPostgres(db),
	}
}
