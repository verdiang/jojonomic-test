package models

// Harga model
type Harga struct {
	ID           string `json:"id,omitempty" gorm:"primary_key"`
	HargaTopup   int64  `json:"harga_topup" gorm:"column:harga_topup"`
	HargaBuyback int64  `json:"harga_buyback" gorm:"column:harga_buyback"`
	AdminID      string `json:"admin_id,omitempty" gorm:"column:admin_id"`
	CreatedAt    int    `json:"created_at,omitempty" gorm:"column:created_at"`
}

// Response model for response
type Response struct {
	Error  bool        `json:"error"`
	ReffID string      `json:"reff_id,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
