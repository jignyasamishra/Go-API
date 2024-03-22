package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]loginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "789GHI",
		Username:  "jason",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *loginDetails {
	//stimulate DB calls
	time.Sleep(time.Second * 1)

	var clientData = loginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil

	}
	return &clientData
}

func (d *mockDB) GetCoinBalance(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil

	}
	return &clientData

}
func (d *mockDB) SetupDatabase() error {

	return nil

}
