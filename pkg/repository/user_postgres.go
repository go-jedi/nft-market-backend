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

func (r *UserPostgres) RegistrationUser(userForm appl_row.UserCreate) (int, error) { // Регистрация пользователя
	var uid string
	userFormJson, _ := json.Marshal(userForm)
	err := r.db.QueryRow("SELECT uid($1)", 8).Scan(&uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err)
	}
	_, err = r.db.Exec("SELECT user_create($1, $2)", userFormJson, uid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_create из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) UpdateLanguage(userFormLang appl_row.UserUpdateLanguage) (int, error) {
	userFormLangJson, _ := json.Marshal(userFormLang)
	_, err := r.db.Exec("SELECT user_update_lang($1)", userFormLangJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_update_lang из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *UserPostgres) CheckIsLanguage(teleId int64) ([]appl_row.CheckUserLanguageResponse, int, error) {
	var userData []appl_row.CheckUserLanguageResponse
	var userDataByte []byte
	err := r.db.QueryRow("SELECT user_check_lang($1)", teleId).Scan(&userDataByte)
	if err != nil {
		return []appl_row.CheckUserLanguageResponse{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_lang из базы данных, %s", err)
	}
	err = json.Unmarshal(userDataByte, &userData)
	if err != nil {
		return []appl_row.CheckUserLanguageResponse{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции CheckIsLanguage, %s", err)
	}
	return userData, http.StatusOK, nil
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
