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
