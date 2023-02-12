package server

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/devansh42/pushNpull/exchange"
)

const maxResponseDelay = 10 // 10 secs
type StockPrice struct {
	Price int `json:"price"`
}

type server struct {
	exchange      exchange.Exchange
	stockPrice    int
	newDataSignal chan struct{}
	newClients    []chan struct{}
	rw, crw       *sync.RWMutex
}

type Server interface {
	PushHandler(wr http.ResponseWriter, r *http.Request)
	PullHandler(wr http.ResponseWriter, r *http.Request)
}

func NewServer() Server {
	serv := server{
		exchange:      exchange.NewExchange(),
		stockPrice:    2000,
		newDataSignal: make(chan struct{}),
		rw:            new(sync.RWMutex),
		crw:           new(sync.RWMutex),
	}
	rand.Seed(time.Now().Unix())
	go serv.startExchange()
	go serv.startBroadCasting()
	return &serv
}

func (s *server) getStockPriceFromDB() int {
	defer s.rw.Unlock()
	s.rw.Lock()

	return s.stockPrice
}

func (s *server) startExchange() {
	for {
		s.rw.Lock()
		s.stockPrice = s.exchange.GetPrice()
		s.rw.Unlock()
		s.newDataSignal <- struct{}{}
		log.Print("Latest Stock Price: ", s.getStockPriceFromDB())
		time.Sleep(time.Second * time.Duration(rand.Intn(maxResponseDelay)))

	}
}

func (s *server) registerNewClient(client chan struct{}) {
	defer s.crw.Unlock()
	s.crw.Lock()

	s.newClients = append(s.newClients, client)
}

func (s *server) removeClient(client chan struct{}) {

	var target = -1
	s.crw.RLock()
	if len(s.newClients) == 0 {
		return
	}
	for i, v := range s.newClients {
		if v == client {
			target = i
			break
		}
	}
	s.crw.RUnlock()

	if target > -1 {
		s.crw.Lock()
		s.newClients = append(s.newClients[:target], s.newClients[target+1:]...)
		s.crw.Unlock()
	}

}
func (s *server) startBroadCasting() {
	for range s.newDataSignal {
		s.crw.RLock()
		for _, ch := range s.newClients {
			ch <- struct{}{}
		}
		s.crw.RUnlock()
	}
}
