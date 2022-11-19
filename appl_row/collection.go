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
