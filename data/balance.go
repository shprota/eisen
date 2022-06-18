package data

type Balance struct {
	Currency string  `json:"currency"`
	Balance  float32 `json:"balance"`
}

type BalancesDto struct {
	Meta     PaginationMeta `json:"meta"`
	Balances []Balance      `json:"balances"`
}
