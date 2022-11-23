package appl_row

type ExchangeRatesGet struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UserCreate struct {
	TeleId      int64  `json:"tele_id"`
	TeleName    string `json:"tele_name"`
	TeleIdAdmin int64  `json:"tele_id_admin"`
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
	Id           int     `json:"id"`
	TeleId       int64   `json:"tele_id"`
	Balance      float64 `json:"balance"`
	IsPremium    bool    `json:"is_premium"`
	Conclusion   float64 `json:"conclusion"`
	Verification bool    `json:"verification"`
}

type UserMinPrice struct {
	MinimPrice float64 `json:"minim_price"`
}

type AdminByUser struct {
	TeleId   int64  `json:"tele_id"`
	TeleName string `json:"tele_name"`
}
