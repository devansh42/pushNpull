package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/devansh42/pushNpull/server"
)

func main() {

	port := flag.Int("port", 4242, "default http server port")

	flag.Parse()

	for !flag.Parsed() {
	}

	serv := server.NewServer()
	http.HandleFunc("/stock/poll", serv.PullHandler)
	http.HandleFunc("/stock/push", serv.PushHandler)
	log.Printf("Starting server at %d", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
