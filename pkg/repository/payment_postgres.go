package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{
		db: db,
	}
}

func (r *PaymentPostgres) CreatePayment(paymentForm appl_row.PaymentCreate) (int, error) {
	paymentFormJson, err := json.Marshal(paymentForm)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка конвертации из struct to json из функции CreatePayment, %s", err)
	}
	_, err = r.db.Exec("SELECT payment_create($1)", paymentFormJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции payment_create из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *PaymentPostgres) GetAllPayments() ([]appl_row.Payment, int, error) {
	var payments []appl_row.Payment
	var paymentsByte []byte
	err := r.db.QueryRow("SELECT json_agg(row_to_json(p.*)) FROM payments p;").Scan(&paymentsByte)
	if err != nil {
		return []appl_row.Payment{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции GetAllPayments из базы данных, %s", err)
	}
	err = json.Unmarshal(paymentsByte, &payments)
	if err != nil {
		return []appl_row.Payment{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetAllPayments, %s", err)
	}
	return payments, http.StatusOK, nil
}
