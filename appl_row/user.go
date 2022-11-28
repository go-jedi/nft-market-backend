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

type UserBalance struct {
	Balance float64 `json:"balance"`
}

type UserBuyToken struct {
	TeleId     int64   `json:"tele_id"`
	TokenUid   string  `json:"token_uid"`
	TokenPrice float64 `json:"token_price"`
}

type UserNft struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Author   string  `json:"author"`
	TokenUid string  `json:"token_uid"`
}

type UserGetNft struct {
	NftBuy  []UserNft `json:"nft_buy"`
	NftSell []UserNft `json:"nft_sell"`
}

type UserSellToken struct {
	TeleId     int64   `json:"tele_id"`
	TokenUid   string  `json:"token_uid"`
	TokenPrice float64 `json:"token_price"`
}

type UserCheckToken struct {
	BuyNft []string `json:"buy_nft"`
}

type UserGetUserPaymentEvent struct {
	TeleId    int64   `json:"tele_id"`
	Uid       string  `json:"uid"`
	NameToken string  `json:"name_token"`
	Price     float64 `json:"price"`
}

type UserWithDrawEventCreate struct {
	TeleId int64   `json:"tele_id"`
	Price  float64 `json:"price"`
}

type WithDrawEventGet struct {
	TeleId     int64   `json:"tele_id"`
	Uid        string  `json:"uid"`
	Price      float64 `json:"price"`
	IsFinished bool    `json:"is_finished"`
}
