package main

import (
	"flag"
	"log"
	"time"
)

func main() {

	dur := flag.Duration("interval", time.Duration(5)*time.Second, "polling interval")
	pull := flag.Bool("pull", false, "uses pull based client")
	serv := flag.String("serv", "http://localhost:4242/stock", "default server url")

	flag.Parse()

	for !flag.Parsed() {
	}

	if *pull {
		log.Print("Starting pull client")
		RunPullingClient(*serv+"/poll", *dur)
	} else {
		log.Print("Starting push client")
		RunPushClient(*serv + "/push")
	}

}
