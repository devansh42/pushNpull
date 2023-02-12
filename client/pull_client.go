package main

import "time"

func RunPullingClient(url string, pollingDuration time.Duration) {
	runClient(url, pollingDuration)
}
