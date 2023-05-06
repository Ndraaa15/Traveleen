package model

type NewExchangeTrash struct {
	Category string  `json:"category"`
	Mass     float64 `json:"mass"`
}

type ValidateCode struct {
	Code      string `json:"code"`
	IsSuccess bool   `json:"is_success"`
}
