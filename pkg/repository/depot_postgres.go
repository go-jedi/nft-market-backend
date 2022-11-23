package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

type DepotPostgres struct {
	db *sqlx.DB
}

func NewDepotPostgres(db *sqlx.DB) *DepotPostgres {
	return &DepotPostgres{
		db: db,
	}
}

func (r *DepotPostgres) CreateDepot(depotForm appl_row.DepotCreate) (int, error) {
	depotFormJson, err := json.Marshal(depotForm)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка конвертации из struct to json из функции CreatePayment, %s", err)
	}
	_, err = r.db.Exec("SELECT depot_create($1)", depotFormJson)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции payment_create из базы данных, %s", err)
	}
	return http.StatusOK, nil
}
