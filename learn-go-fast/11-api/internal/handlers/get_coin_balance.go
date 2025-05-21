package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cankurttekin/gotutorial/11-api/api"
	"github.com/cankurttekin/gotutorial/11-api/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.GetCoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
	
	var tokenDetails *tools.CoinDetails
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.GetCoinBalanceResponse{
		Balance: (*tokenDetails).Coins,	
		Code:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
