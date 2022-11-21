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

func (r *CollectionPostgres) CreateToken(tokenForm appl_row.TokenCreate) (int, error) {
	var uidToken string
	tokenFormJson, err := json.Marshal(tokenForm)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка конвертации из struct to json из функции CreateToken, %s", err)
	}
	err = r.db.QueryRow("SELECT token_uid($1)", 4).Scan(&uidToken)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции token_uid из базы данных, %s", err)
	}
	_, err = r.db.Exec("SELECT token_create($1, $2)", tokenFormJson, uidToken)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции token_create из базы данных, %s", err)
	}
	return http.StatusOK, nil
}

func (r *CollectionPostgres) GetAllTokensCollection(uidCollection string) ([]appl_row.TokensGetByCollection, int, error) {
	var tokensCollection []appl_row.TokensGetByCollection
	var tokensCollectionByte []byte
	err := r.db.QueryRow("SELECT get_tokens_by_collection($1)", uidCollection).Scan(&tokensCollectionByte)
	if err != nil {
		return []appl_row.TokensGetByCollection{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции token_uid из базы данных, %s", err)
	}
	err = json.Unmarshal(tokensCollectionByte, &tokensCollection)
	if err != nil {
		return []appl_row.TokensGetByCollection{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetAllTokensCollection, %s", err)
	}
	return tokensCollection, http.StatusOK, nil
}

func (r *CollectionPostgres) GetToken(tokenUid string) ([]appl_row.Token, int, error) {
	var token []appl_row.Token
	var tokenByte []byte
	err := r.db.QueryRow("SELECT get_token_by_uid($1)", tokenUid).Scan(&tokenByte)
	if err != nil {
		return []appl_row.Token{}, http.StatusInternalServerError, fmt.Errorf("ошибка выполнения функции GetToken из базы данных, %s", err)
	}
	err = json.Unmarshal(tokenByte, &token)
	if err != nil {
		return []appl_row.Token{}, http.StatusInternalServerError, fmt.Errorf("ошибка конвертации в функции GetToken, %s", err)
	}
	return token, http.StatusOK, nil
}
