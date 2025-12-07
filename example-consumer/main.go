package exampleconsumer
package main

import (
	"fmt"
	"time"

	ratelimiter "example.com/policies/rate-limiter"
)

func main() {
	rl := ratelimiter.NewRateLimiter(10, time.Minute)
	fmt.Println("Rate limiter created")
	if rl.Allow() {
		fmt.Println("Request allowed")
	} else {
		fmt.Println("Request denied")
	}
}