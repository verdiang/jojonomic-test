package models

// Price model for price
type Price struct {
	ID           string `json:"id"`
	TopupPrice   int64  `json:"harga_topup"`
	BuybackPrice int64  `json:"harga_buyback"`
	AdminID      string `json:"admin_id"`
}

// Response model for response
type Response struct {
	Error  bool        `json:"error"`
	ReffID string      `json:"reff_id,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
