package appl_row

type UserCreate struct {
	TeleId   int64  `json:"tele_id"`
	TeleName string `json:"tele_name"`
}

type UserUpdateLanguage struct {
	TeleId int64  `json:"tele_id"`
	Lang   string `json:"lang"`
}

type UserLanguage struct {
	Id     int    `json:"id"`
	TeleId int64  `json:"tele_id"`
	Lang   string `json:"lang"`
}

type UserCurrency struct {
	Id       int    `json:"id"`
	TeleId   int64  `json:"tele_id"`
	Currency string `json:"currency"`
}

type UserUpdateCurrency struct {
	TeleId   int64  `json:"tele_id"`
	Currency string `json:"currency"`
}

type UserProfile struct {
	Id           int   `json:"id"`
	TeleId       int64 `json:"tele_id"`
	Balance      int64 `json:"balance"`
	Conclusion   int64 `json:"conclusion"`
	Verification bool  `json:"verification"`
}
