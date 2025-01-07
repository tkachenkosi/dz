package main

import (
	"fmt"
)

func main() {
	const (
		USD_EUR float64 = 0.9667
		USD_RUB float64 = 101.68
	)

	EUR_RUB := USD_RUB / USD_EUR
	fmt.Println("EUR в RUB =", EUR_RUB)
}
