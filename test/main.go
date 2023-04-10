package main

import (
	"fmt"
	"os"
	"ping"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <hostname> [timeout] [count]\n", os.Args[0])
		os.Exit(1)
	}

	hostname := os.Args[1]
	timeout := 1 * time.Second
	count := 4

	if len(os.Args) >= 3 {
		t, err := time.ParseDuration(os.Args[2] + "s")
		if err == nil {
			timeout = t
		}
	}

	if len(os.Args) >= 4 {
		c, err := strconv.Atoi(os.Args[3])
		if err == nil {
			count = c
		}
	}

	err := ping.Ping(hostname, timeout, count)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
