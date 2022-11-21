package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{
		db: db,
	}
}

func (r *AdminPostgres) CheckIsAdmin(teleId int64) (bool, int, error) {
	var isAdmin bool
	err := r.db.QueryRow("SELECT user_check_admin($1)", teleId).Scan(&isAdmin)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_check_admin из базы данных, %s", err)
	}
	return isAdmin, http.StatusOK, nil
}

func (r *AdminPostgres) CreateReferral(referralForm appl_row.ReferralCreate) (int, error) {
	referralFormJson, _ := json.Marshal(referralForm)
	_, err := r.db.Exec("SELECT user_create_referal($1)", referralFormJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_create_referal из базы данных, %s", err)
	}

	return http.StatusOK, nil
}

func (r *AdminPostgres) GetUsersReferral(teleId int64) ([]appl_row.Referral, int, error) {
	var usersReferralAdmin []appl_row.Referral
	var usersReferralAdminByte []byte
	err := r.db.QueryRow("SELECT admin_get_users_ref($1)", teleId).Scan(&usersReferralAdminByte)
	if err != nil {
		return []appl_row.Referral{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_get_users_ref из базы данных, %s", err)
	}
	err = json.Unmarshal(usersReferralAdminByte, &usersReferralAdmin)
	if err != nil {
		return []appl_row.Referral{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUsersReferral, %s", err)
	}

	return usersReferralAdmin, http.StatusOK, nil
}

func (r *AdminPostgres) AdminGetUserProfile(teleId int64) ([]appl_row.AdminUserProfileGet, int, error) {
	var userProfile []appl_row.AdminUserProfileGet
	var userProfileByte []byte
	err := r.db.QueryRow("SELECT admin_get_user_profile($1)", teleId).Scan(&userProfileByte)
	if err != nil {
		return []appl_row.AdminUserProfileGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_get_user_profile из базы данных, %s", err)
	}
	err = json.Unmarshal(userProfileByte, &userProfile)
	if err != nil {
		return []appl_row.AdminUserProfileGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции AdminGetUserProfile, %s", err)
	}
	return userProfile, http.StatusOK, nil
}

func (r *AdminPostgres) UpdatePremium(teleId int64) (int, error) {
	_, err := r.db.Exec("SELECT user_update_premium($1)", teleId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_update_premium из базы данных, %s", err)
	}
	return http.StatusOK, nil
}
