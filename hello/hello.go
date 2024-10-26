package main

import (
	"fmt"
	"time"
)

import "rsc.io/quote"

func main() {
	fmt.Println("Current time:", time.Now())
	fmt.Println(quote.Hello())
}
