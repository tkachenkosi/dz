package main

import (
	"fmt"
)

const (
	ERR_MSG = "Не верный код валюты. Повторите ввод."
)

func main() {
	fmt.Println("КОНВЕРТОР ВАЛЮТ")
	valIn := getValIn()
	summa := getSum()
	valOut := getValOut(valIn)

	calculateCurs(summa, valIn, valOut)
}

func codeVal(codeNum int) string {
	switch codeNum {
	case 1:
		return "USD"
	case 2:
		return "EUR"
	case 3:
		return "RUS"
	default:
		return "ERR"
	}
}

func getValIn() int {
	var val int

	for {
		fmt.Print("Введите исходную валюту 1-USD 2-EUR 3-RUB: ")
		fmt.Scan(&val)

		if val >= 1 && val <= 3 {
			return val
		} else {
			fmt.Println(ERR_MSG)
		}
	}
}

func getValOut(valIn int) int {
	var val int
	var listCodeVal string

	// убираем из списка код исходной валюты
	for i := 1; i < 4; i++ {
		if i != valIn {
			listCodeVal = fmt.Sprintf("%s %d-%s", listCodeVal, i, codeVal(i))
		}
	}

	for {
		fmt.Print("Введите целевую валюту", listCodeVal, ": ")
		fmt.Scan(&val)

		if val < 1 && val > 3 {
			fmt.Println(ERR_MSG)
			continue
		}

		if val == valIn {
			fmt.Println("Исходная и целевая валюта должны быть разными. Повторите ввод.")
		} else {
			return val
		}

	}
}

func getSum() float64 {
	var sum float64

	for {
		fmt.Print("Введите сумму для конвертации: ")
		fmt.Scan(&sum)

		if sum <= 0 {
			fmt.Println("Не верное значение суммы. Повторите ввод.")
		} else {
			return sum
		}
	}
}

func calculateCurs(sumIn float64, valIn, valOut int) {
	// заданы курсы двух базовых пар валют
	const (
		USD_EUR float64 = 0.9667
		USD_RUB float64 = 101.68
	)

	var valCurs float64

	switch {
	case valIn == 1 && valOut == 2:
		valCurs = USD_EUR
	case valIn == 1 && valOut == 3:
		valCurs = USD_RUB
	case valIn == 2 && valOut == 1:
		valCurs = 1.0 / USD_EUR
	case valIn == 2 && valOut == 3:
		valCurs = USD_RUB / USD_EUR
	case valIn == 3 && valOut == 1:
		valCurs = 1.0 / USD_RUB
	case valIn == 3 && valOut == 2:
		valCurs = USD_EUR / USD_RUB
	}

	fmt.Printf("Результат конвертации %.2f %s в %s = %.2f по кросс-курсу %.5f\n",
		sumIn, codeVal(valIn), codeVal(valOut), sumIn*valCurs, valCurs)
}
