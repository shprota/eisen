package data

type ExchangeResultDto struct {
	Success bool    `json:"success"`
	Balance Balance `json:"balance"`
}

type ExchangeDto struct {
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}
