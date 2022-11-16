package models

// Topup model for topup request
type Topup struct {
	ID        string `json:"id"`
	Gram      string `json:"gram"`
	Harga     int    `json:"harga"`
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

// Response model for response
type Response struct {
	Error  bool        `json:"error"`
	ReffID string      `json:"reff_id"`
	Data   interface{} `json:"data,omitempty"`
}

// Harga model for price
type Harga struct {
	ID           string `json:"id" gorm:"primary_key"`
	HargaTopup   int64  `json:"harga_topup" gorm:"column:harga_topup"`
	HargaBuyback int64  `json:"harga_buyback" gorm:"column:harga_buyback"`
	AdminID      string `json:"admin_id" gorm:"column:admin_id"`
	CreatedAt    int    `json:"created_at" gorm:"column:created_at"`
}
