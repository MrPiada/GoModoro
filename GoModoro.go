package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Time at start:", start)
	for {
		elapsed := time.Now().Sub(start).Seconds()
		fmt.Println("Time since start:", elapsed)
	}
}
