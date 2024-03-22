package api

// coin balance params
type CoinBalanceParams struct {
	Username string
}

// coin balance response
type CoinBalanceResponse struct {
	Code    int
	Balance int64
}
type Error struct {
	Code    int
	Message string
}
