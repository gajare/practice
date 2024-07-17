package main

import (
	"fmt"
	"time"
)

func main() {
	TrackTime(time.Now())
	time.Sleep(5 * time.Second)
}

func TrackTime(start time.Time) time.Duration {
	elapsed := time.Since(start)
	fmt.Println("elapsed :", elapsed.String())
	return elapsed
}
