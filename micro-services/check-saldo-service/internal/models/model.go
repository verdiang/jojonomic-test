package models

// Rekening model for rekening
type Rekening struct {
	ID    string  `json:"id"`
	Norek string  `json:"norek"`
	Saldo float32 `json:"saldo"`
}

// Response model for response
type Response struct {
	Error  bool        `json:"error"`
	ReffID string      `json:"reff_id"`
	Data   interface{} `json:"data,omitempty"`
}
