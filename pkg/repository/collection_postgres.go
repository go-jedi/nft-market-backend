package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

type CollectionPostgres struct {
	db *sqlx.DB
}

func NewCollectionPostgres(db *sqlx.DB) *CollectionPostgres {
	return &CollectionPostgres{
		db: db,
	}
}

func (r *CollectionPostgres) CreateCollection(collectionForm appl_row.CollectionCreate) (int, error) {
	var uidCollection string
	collectionFormJson, err := json.Marshal(collectionForm)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка конвертации из struct to json из функции CreateCollection, %s", err)
	}
	err = r.db.QueryRow("SELECT collection_uid($1)", 4).Scan(&uidCollection)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции collection_uid из базы данных, %s", err)
	}
	_, err = r.db.Exec("SELECT collection_create($1, $2)", collectionFormJson, uidCollection)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции payment_create из базы данных, %s", err)
	}

	return http.StatusOK, nil
}

func (r *CollectionPostgres) GetAllCollections() ([]appl_row.Collection, int, error) {
	var collections []appl_row.Collection
	var collectionsByte []byte
	err := r.db.QueryRow("SELECT json_agg(row_to_json(p.*)) FROM collections p;").Scan(&collectionsByte)
	if err != nil {
		return []appl_row.Collection{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции GetAllCollections из базы данных, %s", err)
	}
	err = json.Unmarshal(collectionsByte, &collections)
	if err != nil {
		return []appl_row.Collection{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetAllPayments, %s", err)
	}
	return collections, http.StatusOK, nil
}
