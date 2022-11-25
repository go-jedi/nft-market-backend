package appl_row

type DepotCreate struct {
	MammothId       int64   `json:"mammoth_id"`
	MammothUsername string  `json:"mammoth_username"`
	WorkerId        int64   `json:"worker_id"`
	WorkerUsername  string  `json:"worker_username"`
	Amount          float64 `json:"amount"`
	IsShowName      bool    `json:"is_show_name"`
}
