package appl_row

type ReferralCreate struct {
	TeleId        int64  `json:"tele_id"`
	TeleName      string `json:"tele_name"`
	AdminReferral int64  `json:"admin_referral"`
}

type Referral struct {
	Id            int64  `json:"id"`
	TeleId        int64  `json:"tele_id"`
	TeleName      string `json:"tele_name"`
	Created       string `json:"created"`
	AdminReferral int64  `json:"admin_referral"`
}

type AdminUserProfileGet struct {
	Id           int64  `json:"id"`
	TeleId       int64  `json:"tele_id"`
	TeleName     string `json:"tele_name"`
	Balance      int64  `json:"balance"`
	IsPremium    bool   `json:"is_premium"`
	Verification bool   `json:"verification"`
	Conclusion   int64  `json:"conclusion"`
}
