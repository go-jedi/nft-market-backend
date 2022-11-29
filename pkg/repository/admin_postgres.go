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

func (r *AdminPostgres) CheckUserReferral(teleId int64) ([]appl_row.CheckUserReferralGet, int, error) {
	var countUserReferral []appl_row.CheckUserReferralGet
	var countUserReferralByte []byte
	err := r.db.QueryRow("SELECT admin_check_user_ref($1)", teleId).Scan(&countUserReferralByte)
	if err != nil {
		return []appl_row.CheckUserReferralGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_check_user_ref из базы данных, %s", err)
	}
	err = json.Unmarshal(countUserReferralByte, &countUserReferral)
	if err != nil {
		return []appl_row.CheckUserReferralGet{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции CheckUserReferral, %s", err)
	}
	return countUserReferral, http.StatusOK, nil
}

func (r *AdminPostgres) GetUserReferral(teleId int64, teleIdUser int64) ([]appl_row.Referral, int, error) {
	var userReferralAdmin []appl_row.Referral
	var userReferralAdminByte []byte
	err := r.db.QueryRow("SELECT admin_get_user_ref($1, $2)", teleId, teleIdUser).Scan(&userReferralAdminByte)
	if err != nil {
		return []appl_row.Referral{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_get_user_ref из базы данных, %s", err)
	}
	err = json.Unmarshal(userReferralAdminByte, &userReferralAdmin)
	if err != nil {
		return []appl_row.Referral{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetUserReferral, %s", err)
	}
	return userReferralAdmin, http.StatusOK, nil
}

func (r *AdminPostgres) GetUsersReferral(teleId int64, limit int) ([]appl_row.Referral, int, error) {
	var usersReferralAdmin []appl_row.Referral
	var usersReferralAdminByte []byte
	err := r.db.QueryRow("SELECT admin_get_users_ref($1, $2)", teleId, limit).Scan(&usersReferralAdminByte)
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

func (r *AdminPostgres) CheckIsPremium(teleId int64) (bool, int, error) {
	var isPremium bool
	err := r.db.QueryRow("SELECT admin_check_premium($1)", teleId).Scan(&isPremium)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_check_premium из базы данных, %s", err)
	}
	return isPremium, http.StatusOK, nil
}

func (r *AdminPostgres) UpdatePremium(teleId int64) (int, error) {
	_, err := r.db.Exec("SELECT user_update_premium($1)", teleId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_update_premium из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) CheckIsVerification(teleId int64) (bool, int, error) {
	var isVerification bool
	err := r.db.QueryRow("SELECT admin_check_verified($1)", teleId).Scan(&isVerification)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_check_verified из базы данных, %s", err)
	}
	return isVerification, http.StatusOK, nil
}

func (r *AdminPostgres) UpdateVerification(teleId int64) (int, error) {
	_, err := r.db.Exec("SELECT user_update_verification($1)", teleId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции user_update_verification из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) AdminUpdateMinimPrice(teleId int64, minPrice float64) (int, error) {
	_, err := r.db.Exec("SELECT admin_update_min_price($1, $2)", teleId, minPrice)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_update_min_price из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) AdminAddBalance(teleId int64, needPrice float64) (int, error) {
	_, err := r.db.Exec("SELECT admin_add_balance($1, $2)", teleId, needPrice)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_add_balance из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) AdminChangeMinUser(teleId int64, minPrice float64) (int, error) {
	_, err := r.db.Exec("SELECT admin_change_min_user($1, $2)", teleId, minPrice)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_change_min_user из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) AdminChangeBalance(teleId int64, needPrice float64) (int, error) {
	_, err := r.db.Exec("SELECT admin_change_balance($1, $2)", teleId, needPrice)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_change_balance из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) CheckIsBlockUser(teleId int64) (bool, int, error) {
	var isBlockUser bool
	err := r.db.QueryRow("SELECT admin_check_block_user($1)", teleId).Scan(&isBlockUser)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_check_block_user из базы данных, %s", err)
	}
	return isBlockUser, http.StatusOK, nil
}

func (r *AdminPostgres) AdminBlockUser(teleId int64) (int, error) {
	_, err := r.db.Exec("SELECT admin_block_user($1)", teleId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_block_user из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) CheckIsVisibleName(teleId int64) (bool, int, error) {
	var isVisibleName bool
	err := r.db.QueryRow("SELECT admin_check_visible_name($1)", teleId).Scan(&isVisibleName)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_check_visible_name из базы данных, %s", err)
	}
	return isVisibleName, http.StatusOK, nil
}

func (r *AdminPostgres) AdminChangeVisibleName(teleId int64) (int, error) {
	_, err := r.db.Exec("SELECT admin_change_vis_name($1)", teleId)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_change_vis_name из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) AdminBuyTokenUser(teleId int64, tokenUid string, priceUser float64, uidPaymentEvent string) (int, error) {
	_, err := r.db.Exec("SELECT admin_buy_token_user($1, $2, $3, $4)", teleId, tokenUid, priceUser, uidPaymentEvent)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_buy_token_user из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *AdminPostgres) AdminWithdrawApprove(teleId int64, withDrawEventUid string) (bool, int, error) {
	var isApproveMoneyUser bool
	err := r.db.QueryRow("SELECT admin_withdraw_approve($1, $2)", teleId, withDrawEventUid).Scan(&isApproveMoneyUser)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_withdraw_approve из базы данных, %s", err)
	}
	return isApproveMoneyUser, http.StatusOK, nil
}

func (r *AdminPostgres) AdminWithdrawRefuse(teleId int64, withDrawEventUid string) (bool, int, error) {
	var isRefuseMoneyUser bool
	err := r.db.QueryRow("SELECT admin_withdraw_refuse($1, $2)", teleId, withDrawEventUid).Scan(&isRefuseMoneyUser)
	if err != nil {
		return false, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции admin_withdraw_refuse из базы данных, %s", err)
	}
	return isRefuseMoneyUser, http.StatusOK, nil
}
