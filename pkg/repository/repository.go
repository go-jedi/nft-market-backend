package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

type TodoUser interface {
	GetAllExchangeRates() ([]appl_row.ExchangeRatesGet, int, error)
	CheckAuth(teleId int64) (bool, int, error)
	RegistrationUser(userForm appl_row.UserCreate) (int, error)
	GetUserLanguage(teleId int64) ([]appl_row.UserLanguage, int, error)
	UpdateLanguage(userForm appl_row.UserUpdateLanguage) (int, error)
	GetUserCurrency(teleId int64) ([]appl_row.UserCurrency, int, error)
	UpdateCurrency(userFormCurrency appl_row.UserUpdateCurrency) (int, error)
	CheckIsTerms(teleId int64) (bool, int, error)
	AgreeTerms(teleId int64) (int, error)
	GetUserProfile(teleId int64) ([]appl_row.UserProfile, int, error)
	GetUserMinPrice(teleId int64) ([]appl_row.UserMinPrice, int, error)
	GetAdminByUser(teleId int64) ([]appl_row.AdminByUser, int, error)
	GetUserBalance(teleId int64) ([]appl_row.UserBalance, int, error)
	CheckUserToken(teleId int64, tokenUid string) (bool, int, error)
	BuyUserToken(userBuyTokenForm appl_row.UserBuyToken) (int, error)
	SellUserToken(userSellTokenForm appl_row.UserSellToken) (string, int, error)
	GetUserNft(teleId int64) (appl_row.UserGetNft, int, error)
	GetUserPaymentEvent(eventUid string) ([]appl_row.UserGetUserPaymentEvent, int, error)
	CreateWithDrawEvent(userWithDrawForm appl_row.UserWithDrawEventCreate) (string, int, error)
	GetWithDrawEvent(withDrawEventUid string) ([]appl_row.WithDrawEventGet, int, error)
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
	CheckUserReferral(teleId int64) ([]appl_row.CheckUserReferralGet, int, error)
	GetUserReferral(teleId int64, teleIdUser int64) ([]appl_row.Referral, int, error)
	GetUsersReferral(teleId int64, limit int) ([]appl_row.Referral, int, error)
	AdminGetUserProfile(teleId int64) ([]appl_row.AdminUserProfileGet, int, error)
	CheckIsPremium(teleId int64) (bool, int, error)
	UpdatePremium(teleId int64) (int, error)
	AdminUpdateMinimPrice(teleId int64, minPrice float64) (int, error)
	CheckIsVerification(teleId int64) (bool, int, error)
	UpdateVerification(teleId int64) (int, error)
	AdminAddBalance(teleId int64, needPrice float64) (int, error)
	AdminChangeMinUser(teleId int64, minPrice float64) (int, error)
	AdminChangeBalance(teleId int64, needPrice float64) (int, error)
	CheckIsBlockUser(teleId int64) (bool, int, error)
	AdminBlockUser(teleId int64) (int, error)
	CheckIsVisibleName(teleId int64) (bool, int, error)
	AdminChangeVisibleName(teleId int64) (int, error)
	AdminBuyTokenUser(teleId int64, tokenUid string, priceUser float64, uidPaymentEvent string) (int, error)
	AdminWithdrawApprove(teleId int64, withDrawEventUid string) (bool, int, error)
}

type TodoDepot interface {
	CreateDepot(depotForm appl_row.DepotCreate) (int, error)
}

type Repository struct {
	TodoUser
	TodoPayment
	TodoCollection
	TodoAdmin
	TodoDepot
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoUser:       NewUserPostgres(db),
		TodoPayment:    NewPaymentPostgres(db),
		TodoCollection: NewCollectionPostgres(db),
		TodoAdmin:      NewAdminPostgres(db),
		TodoDepot:      NewDepotPostgres(db),
	}
}
