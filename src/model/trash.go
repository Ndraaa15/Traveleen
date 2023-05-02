package model

type NewExchangeTrash struct {
	Category string  `json:"category"`
	Location string  `json:"location"`
	Mass     float64 `json:"mass"`
}
