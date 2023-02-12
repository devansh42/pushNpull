package exchange

import (
	"math/rand"
	"time"
)

type exchange struct{}
type Exchange interface {
	GetPrice() int
}

func NewExchange() Exchange {
	rand.Seed(time.Now().Unix())
	return exchange{}
}

func (e exchange) GetPrice() int {
	return 2000 + rand.Intn(200)
}
