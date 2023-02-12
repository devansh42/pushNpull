package server

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/devansh42/pushNpull/exchange"
)

const maxResponseDelay = 10 // 10 secs
type StockPrice struct {
	Price int `json:"price"`
}

type server struct {
	exchange   exchange.Exchange
	stockPrice int
	priceCh    chan int
}

type Server interface {
	PushHandler(wr http.ResponseWriter, r *http.Request)
	PullHandler(wr http.ResponseWriter, r *http.Request)
}

func NewServer() Server {
	serv := server{
		exchange:   exchange.NewExchange(),
		stockPrice: 2000,
		priceCh:    make(chan int),
	}
	rand.Seed(time.Now().Unix())
	go serv.startExchange()
	return &serv
}

func (s *server) getStockPriceFromDB() int {
	return s.stockPrice
}

func (s *server) getOnlyLatestPriceFromDB() int {
	s.stockPrice = <-s.priceCh
	return s.getStockPriceFromDB()
}

func (s *server) startExchange() {
	for {
		time.Sleep(time.Second * time.Duration(rand.Intn(maxResponseDelay)))
		s.priceCh <- s.exchange.GetPrice()
	}
}
