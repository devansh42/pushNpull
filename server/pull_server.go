package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) PullHandler(wr http.ResponseWriter, r *http.Request) {
	log.Print("New Pull Connection")
	wr.WriteHeader(http.StatusOK)
	json.NewEncoder(wr).Encode(&StockPrice{
		Price: s.getStockPriceFromDB(),
	})
}
