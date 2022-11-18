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

func (r *UserPostgres) RegistrationUser(userForm appl_row.UserCreate) (error, int) { // Регистрация пользователя
	var uid string
	userFormJson, _ := json.Marshal(userForm)
	err := r.db.QueryRow("SELECT uid($1)", 8).Scan(&uid)
	if err != nil {
		return fmt.Errorf("ошибка выполнения функции uid из базы данных, %s", err), http.StatusInternalServerError
	}
	_, err = r.db.Exec("SELECT user_create($1, $2)", userFormJson, uid)
	if err != nil {
		return fmt.Errorf("ошибка выполнения функции user_create из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *UserPostgres) CheckAuth(teleId int64) (bool, error, int) {
	var isAuth bool
	err := r.db.QueryRow("SELECT user_check_auth($1)", teleId).Scan(&isAuth)
	if err != nil {
		return false, fmt.Errorf("ошибка выполнения функции user_check_auth из базы данных, %s", err), http.StatusInternalServerError
	}
	return isAuth, nil, http.StatusOK
}

func (r *UserPostgres) UpdateLanguage(userFormLang appl_row.UserUpdateLanguage) (error, int) {
	userFormLangJson, _ := json.Marshal(userFormLang)
	_, err := r.db.Exec("SELECT user_update_lang($1)", userFormLangJson)
	if err != nil {
		return fmt.Errorf("ошибка выполнения функции user_update_lang из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *UserPostgres) CheckIsLanguage(teleId int64) ([]appl_row.CheckUserLanguageResponse, error, int) {
	var userData []appl_row.CheckUserLanguageResponse
	var userDataByte []byte
	err := r.db.QueryRow("SELECT user_check_lang($1)", teleId).Scan(&userDataByte)
	if err != nil {
		return []appl_row.CheckUserLanguageResponse{}, fmt.Errorf("ошибка выполнения функции user_check_lang из базы данных, %s", err), http.StatusInternalServerError
	}
	err = json.Unmarshal(userDataByte, &userData)
	if err != nil {
		return []appl_row.CheckUserLanguageResponse{}, fmt.Errorf("ошибка конвертации в функции CheckIsLanguage, %s", err), http.StatusInternalServerError
	}
	return userData, nil, http.StatusOK
}

func (r *UserPostgres) UpdateCurrency(userFormCurrency appl_row.UserUpdateCurrency) (error, int) {
	userFormCurrencyJson, _ := json.Marshal(userFormCurrency)
	_, err := r.db.Exec("SELECT user_update_currency($1)", userFormCurrencyJson)
	if err != nil {
		return fmt.Errorf("ошибка выполнения функции user_update_currency из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (r *UserPostgres) CheckIsTerms(teleId int64) (bool, error, int) {
	var isTerms bool
	err := r.db.QueryRow("SELECT user_check_terms($1)", teleId).Scan(&isTerms)
	if err != nil {
		return false, fmt.Errorf("ошибка выполнения функции user_check_terms из базы данных, %s", err), http.StatusInternalServerError
	}
	return isTerms, nil, http.StatusOK
}

func (r *UserPostgres) AgreeTerms(teleId int64) (error, int) {
	_, err := r.db.Exec("SELECT user_agree_terms($1)", teleId)
	if err != nil {
		return fmt.Errorf("ошибка выполнения функции user_agree_terms из базы данных, %s", err), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}
