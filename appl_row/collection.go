package appl_row

type Collection struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Count         int    `json:"count"`
	CollectionUid string `json:"collection_uid"`
}

type CollectionCreate struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Token struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Author         string  `json:"author"`
	Blockchain     string  `json:"blockchain"`
	UidCollection  string  `json:"uid_collection"`
	NameCollection string  `json:"name_collection"`
	TokenUid       string  `json:"token_uid"`
}

type TokenCreate struct {
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Author        string  `json:"author"`
	Blockchain    string  `json:"blockchain"`
	UidCollection string  `json:"uid_collection"`
}

type TokensGetByCollection struct {
	Id             int     `json:"id"`
	NameCollection string  `json:"name_collection"`
	Count          int     `json:"count"`
	NameToken      string  `json:"name_token"`
	PriceToken     float64 `json:"price_token"`
	TokenUid       string  `json:"token_uid"`
}
