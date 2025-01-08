package main

import (
	"fmt"
)

func main() {
	// const (
	// 	USD_EUR float64 = 0.9667
	// 	USD_RUB float64 = 101.68
	// )

	USD_EUR, USD_RUB := getUserInput()

	EUR_RUB := USD_RUB / USD_EUR
	fmt.Println("EUR в RUB =", EUR_RUB)
}

func getUserInput() (float64, float64) {
	var usd_eur float64
	var usd_rub float64

	fmt.Print("Введите курс для USD EUR: ")
	fmt.Scan(&usd_eur)
	fmt.Print("Введите курс для USD RUB: ")
	fmt.Scan(&usd_rub)

	return usd_eur, usd_rub
}

func calculateCurs(sum float64, val1, val2 int) float64 {
	fmt.Println(val1, val2)
	return sum
}
