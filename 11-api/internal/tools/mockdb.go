package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"can": {
		AuthToken: "1234567890",
		Username:  "can",
	},
	"john": {
		AuthToken: "0987654321",
		Username:  "john",
	},
	"jane": {
		AuthToken: "1122334455",
		Username:  "jane",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"can": {
		Coins:    100,
		Username: "can",
	},
	"john": {
		Coins:    200,
		Username: "john",
	},
	"jane": {
		Coins:    300,
		Username: "jane",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate db call
	time.Sleep(1 * time.Millisecond)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}


func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate db call
	time.Sleep(1 * time.Millisecond)

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
















