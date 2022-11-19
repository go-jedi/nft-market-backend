package appl_row

type Payment struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PaymentCreate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
