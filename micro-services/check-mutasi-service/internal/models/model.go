package models

// Topup model for topup request
type Topup struct {
	ID        string `json:"id"`
	Gram      string `json:"gram"`
	Harga     int    `json:"harga_topup"`
	Norek     string `json:"norek"`
	Buyback   int    `json:"harga_buyback"`
	Type      string `json:"type"`
	CreatedAt int    `json:"created_at"`
}

// Rekening model for rekening
type Rekening struct {
	ID    string  `json:"id"`
	Norek string  `json:"norek"`
	Saldo float64 `json:"saldo"`
}

type MutationRequest struct {
	Norek     string `json:"norek"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

// Response model for response
type Response struct {
	Error  bool        `json:"error"`
	ReffID string      `json:"reff_id"`
	Data   interface{} `json:"data,omitempty"`
}
