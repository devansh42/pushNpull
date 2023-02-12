package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/devansh42/pushNpull/server"
)

func runClient(url string, pollingInterval time.Duration) {
	var stockData server.StockPrice

	var nilBody io.Reader
	req, err := http.NewRequest(http.MethodGet, url, nilBody)
	if err != nil {
		log.Fatalf("couldn't create http requets due to %v", err)
	}

	for {
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Printf("couldn't make http request due to %v", err)
			break
		}

		respBody := resp.Body
		jdec := json.NewDecoder(respBody)
		for {
			err = jdec.Decode(&stockData)
			if err != nil {
				if err == io.ErrUnexpectedEOF || err == io.EOF {
					log.Print("Server Closed the connection")
				} else {
					log.Printf("No Data Left to read %v", err)
				}
				break
			}
			log.Printf("Resp: %+v", stockData)
		}
		respBody.Close()

		if pollingInterval == 0 { // It means it's a push client
			break
		}

		time.Sleep(pollingInterval) // Sleeps for polling interval

	}
}
