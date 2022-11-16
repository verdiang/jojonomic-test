package models

// Harga model for table harga
type Harga struct {
	ID           string `json:"id" gorm:"primary_key"`
	HargaTopup   int64  `json:"harga_topup" gorm:"column:harga_topup"`
	HargaBuyback int64  `json:"harga_buyback" gorm:"column:harga_buyback"`
	AdminID      string `json:"admin_id" gorm:"column:admin_id"`
	CreatedAt    int    `json:"created_at" gorm:"column:created_at"`
}

// Response model for response
type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data,omitempty"`
}
