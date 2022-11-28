package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) GetAllExchangeRates() ([]appl_row.ExchangeRatesGet, int, error) {
	var exchangeRates []appl_row.ExchangeRatesGet
	var exchangeRatesByte []byte
	err := r.db.QueryRow("SELECT COALESCE(json_agg(row_to_json(er.*)), '[]') FROM exchange_rates er;").Scan(&exchangeRatesByte)
	if err != nil {
		return []appl_row.ExchangeRatesGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения запроса из базы данных, %s", err)
	}
	err = json.Unmarshal(exchangeRatesByte, &exchangeRates)
	if err != nil {
		return []appl_row.ExchangeRatesGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции ExchangeRates, %s", err)
	}
	return exchangeRates, http.StatusOK, nil
}

func (r *UserPostgres) CheckAuth(teleId int64) (bool, int, error) {
	var isAuth bool
	err := r.db.QueryRow("SELECT user_check_auth($1)", teleId).Scan(&isAuth)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_auth из базы данных, %s", err)
	}
	return isAuth, http.StatusOK, nil
}

func (r *UserPostgres) RegistrationUser(userForm appl_row.UserCreate) (int, error) {
	var uid string
	var isRegisterRes bool
	userFormJson, _ := json.Marshal(userForm)
	err := r.db.QueryRow("SELECT uid($1)", 8).Scan(&uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
	}
	err = r.db.QueryRow("SELECT user_create($1, $2)", userFormJson, uid).Scan(&isRegisterRes)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_create из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) GetUserLanguage(teleId int64) ([]appl_row.UserLanguage, int, error) {
	var userDataLang []appl_row.UserLanguage
	var userDataByte []byte
	err := r.db.QueryRow("SELECT user_get_lang($1)", teleId).Scan(&userDataByte)
	if err != nil {
		return []appl_row.UserLanguage{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_lang из базы данных, %s", err)
	}
	err = json.Unmarshal(userDataByte, &userDataLang)
	if err != nil {
		return []appl_row.UserLanguage{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserLanguage, %s", err)
	}
	return userDataLang, http.StatusOK, nil
}

func (r *UserPostgres) UpdateLanguage(userFormLang appl_row.UserUpdateLanguage) (int, error) {
	userFormLangJson, _ := json.Marshal(userFormLang)
	_, err := r.db.Exec("SELECT user_update_lang($1)", userFormLangJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_update_lang из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) GetUserCurrency(teleId int64) ([]appl_row.UserCurrency, int, error) {
	var userDataCurrency []appl_row.UserCurrency
	var userDataCurrencyByte []byte
	err := r.db.QueryRow("SELECT user_get_currency($1)", teleId).Scan(&userDataCurrencyByte)
	if err != nil {
		return []appl_row.UserCurrency{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_currency из базы данных, %s", err)
	}
	err = json.Unmarshal(userDataCurrencyByte, &userDataCurrency)
	if err != nil {
		return []appl_row.UserCurrency{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserCurrency, %s", err)
	}
	return userDataCurrency, http.StatusOK, nil
}

func (r *UserPostgres) UpdateCurrency(userFormCurrency appl_row.UserUpdateCurrency) (int, error) {
	userFormCurrencyJson, _ := json.Marshal(userFormCurrency)
	_, err := r.db.Exec("SELECT user_update_currency($1)", userFormCurrencyJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_update_currency из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) CheckIsTerms(teleId int64) (bool, int, error) {
	var isTerms bool
	err := r.db.QueryRow("SELECT user_check_terms($1)", teleId).Scan(&isTerms)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_terms из базы данных, %s", err)
	}
	return isTerms, http.StatusOK, nil
}

func (r *UserPostgres) AgreeTerms(teleId int64) (int, error) {
	_, err := r.db.Exec("SELECT user_agree_terms($1)", teleId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_agree_terms из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) GetUserProfile(teleId int64) ([]appl_row.UserProfile, int, error) {
	var userProfile []appl_row.UserProfile
	var userProfileByte []byte
	err := r.db.QueryRow("SELECT user_get_profile($1)", teleId).Scan(&userProfileByte)
	if err != nil {
		return []appl_row.UserProfile{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_profile из базы данных, %s", err)
	}
	err = json.Unmarshal(userProfileByte, &userProfile)
	if err != nil {
		return []appl_row.UserProfile{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserProfile, %s", err)
	}
	return userProfile, http.StatusOK, nil
}

func (r *UserPostgres) GetUserMinPrice(teleId int64) ([]appl_row.UserMinPrice, int, error) {
	var userMinPrice []appl_row.UserMinPrice
	var userMinPriceByte []byte
	err := r.db.QueryRow("SELECT user_get_min_price($1)", teleId).Scan(&userMinPriceByte)
	if err != nil {
		return []appl_row.UserMinPrice{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_min_price из базы данных, %s", err)
	}
	err = json.Unmarshal(userMinPriceByte, &userMinPrice)
	if err != nil {
		return []appl_row.UserMinPrice{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserMinPrice, %s", err)
	}
	return userMinPrice, http.StatusOK, nil
}

func (r *UserPostgres) GetAdminByUser(teleId int64) ([]appl_row.AdminByUser, int, error) {
	var adminByUser []appl_row.AdminByUser
	var adminByUserByte []byte
	err := r.db.QueryRow("SELECT user_get_admin_by_user($1)", teleId).Scan(&adminByUserByte)
	if err != nil {
		return []appl_row.AdminByUser{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_admin_by_user из базы данных, %s", err)
	}
	err = json.Unmarshal(adminByUserByte, &adminByUser)
	if err != nil {
		return []appl_row.AdminByUser{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetAdminByUser, %s", err)
	}
	return adminByUser, http.StatusOK, nil
}

func (r *UserPostgres) GetUserBalance(teleId int64) ([]appl_row.UserBalance, int, error) {
	var userBalance []appl_row.UserBalance
	var userBalanceByte []byte
	err := r.db.QueryRow("SELECT user_get_balance($1)", teleId).Scan(&userBalanceByte)
	if err != nil {
		return []appl_row.UserBalance{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_balance из базы данных, %s", err)
	}
	err = json.Unmarshal(userBalanceByte, &userBalance)
	if err != nil {
		return []appl_row.UserBalance{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции CheckUserBalance, %s", err)
	}
	return userBalance, http.StatusOK, nil
}

func (r *UserPostgres) CheckUserToken(teleId int64, tokenUid string) (bool, int, error) {
	var isHaveToken bool = false
	var tokensUidUser []appl_row.UserCheckToken
	var tokenByte []byte
	err := r.db.QueryRow("SELECT user_check_token($1, $2)", teleId, tokenUid).Scan(&tokenByte)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_token из базы данных, %s", err)
	}
	err = json.Unmarshal(tokenByte, &tokensUidUser)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции CheckUserToken, %s", err)
	}
	for _, value := range tokensUidUser[0].BuyNft {
		if value == tokenUid {
			isHaveToken = true
		}
	}

	return isHaveToken, http.StatusOK, nil
}

func (r *UserPostgres) BuyUserToken(userBuyTokenForm appl_row.UserBuyToken) (int, error) {
	userBuyTokenFormJson, _ := json.Marshal(userBuyTokenForm)
	_, err := r.db.Exec("SELECT user_buy_token($1)", userBuyTokenFormJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_buy_token из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) SellUserToken(userSellTokenForm appl_row.UserSellToken) (string, int, error) {
	var uid string
	var uidCreated string = ""
	userSellTokenFormJson, _ := json.Marshal(userSellTokenForm)
	err := r.db.QueryRow("SELECT payment_event_uid($1)", 8).Scan(&uid)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции payment_event_uid из базы данных, %s", err)
	}
	err = r.db.QueryRow("SELECT user_sell_token($1, $2)", userSellTokenFormJson, uid).Scan(&uidCreated)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_sell_token из базы данных, %s", err)
	}
	return uidCreated, http.StatusOK, nil
}

func (r *UserPostgres) GetUserNft(teleId int64) (appl_row.UserGetNft, int, error) {
	var userNft appl_row.UserGetNft
	var userNftByte []byte
	err := r.db.QueryRow("SELECT user_get_nft($1)", teleId).Scan(&userNftByte)
	if err != nil {
		return appl_row.UserGetNft{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_nft из базы данных, %s", err)
	}
	err = json.Unmarshal(userNftByte, &userNft)
	if err != nil {
		return appl_row.UserGetNft{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserNft, %s", err)
	}
	return userNft, http.StatusOK, nil
}

func (r *UserPostgres) GetUserPaymentEvent(eventUid string) ([]appl_row.UserGetUserPaymentEvent, int, error) {
	var paymentEvent []appl_row.UserGetUserPaymentEvent
	var paymentEventByte []byte
	err := r.db.QueryRow("SELECT user_get_payment_event($1)", eventUid).Scan(&paymentEventByte)
	if err != nil {
		return []appl_row.UserGetUserPaymentEvent{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_payment_event из базы данных, %s", err)
	}
	err = json.Unmarshal(paymentEventByte, &paymentEvent)
	if err != nil {
		return []appl_row.UserGetUserPaymentEvent{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserPaymentEvent, %s", err)
	}
	return paymentEvent, http.StatusOK, nil
}

func (r *UserPostgres) CreateWithDrawEvent(userWithDrawForm appl_row.UserWithDrawEventCreate) (string, int, error) {
	var uid string
	var uidCreated string = ""
	userWithDrawFormJson, _ := json.Marshal(userWithDrawForm)
	err := r.db.QueryRow("SELECT withdraw_event_uid($1)", 4).Scan(&uid)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции withdraw_event_uid из базы данных, %s", err)
	}
	err = r.db.QueryRow("SELECT user_create_withdraw_event($1, $2)", userWithDrawFormJson, uid).Scan(&uidCreated)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_create_withdraw_event из базы данных, %s", err)
	}
	return uidCreated, http.StatusOK, nil
}

func (r *UserPostgres) GetWithDrawEvent(withDrawEventUid string) ([]appl_row.WithDrawEventGet, int, error) {
	var userWithDrawEvent []appl_row.WithDrawEventGet
	var userWithDrawEventByte []byte
	err := r.db.QueryRow("SELECT user_get_withdraw_event($1)", withDrawEventUid).Scan(&userWithDrawEventByte)
	if err != nil {
		return []appl_row.WithDrawEventGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_get_withdraw_event из базы данных, %s", err)
	}
	err = json.Unmarshal(userWithDrawEventByte, &userWithDrawEvent)
	if err != nil {
		return []appl_row.WithDrawEventGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetWithDrawEvent, %s", err)
	}
	return userWithDrawEvent, http.StatusOK, nil
}
