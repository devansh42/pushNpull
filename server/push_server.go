package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) PushHandler(wr http.ResponseWriter, r *http.Request) {
	log.Print("New Push Connection")
	flusher, ok := wr.(http.Flusher)
	if !ok {
		wr.WriteHeader(http.StatusInternalServerError)
		wr.Write([]byte("Streaming un-supported!"))
		return
	}
	wr.Header().Set("Content-Type", "text/event-stream")
	wr.Header().Set("Cache-Control", "no-cache")
	wr.Header().Set("Connection", "keep-alive")
	wr.WriteHeader(http.StatusOK)

	jwr := json.NewEncoder(wr)

	var err error
	jwr.Encode(&StockPrice{ // Initial response
		Price: s.getStockPriceFromDB(),
	})
	var ch = make(chan struct{})
	s.registerNewClient(ch)
	for range ch {
		err = jwr.Encode(&StockPrice{
			Price: s.getStockPriceFromDB(),
		})
		if err != nil {
			log.Printf("closing connection due to %v", err)
			s.removeClient(ch)
			break
		}
		flusher.Flush()
	}
}
