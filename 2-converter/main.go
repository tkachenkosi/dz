package main

import (
	"fmt"
)

const (
	ERR_MSG    = "Не верный код валюты. Повторите ввод."
	COUNT_VALS = 5 // всего валют в списке
	VAL_USD    = 0
)

var codeVal = [5]string{"USD", "EUR", "RUB", "KGS", "CNY"}

func main() {
	fmt.Println("КОНВЕРТОР ВАЛЮТ")
	fromCurrency := getValIn()
	summa := getSum()
	toCurrency := getValOut(fromCurrency)

	calculateCurs(summa, fromCurrency, toCurrency)
}

func getCodeVal(codeNum int) string {
	return codeVal[codeNum]
}

func getCodeVal2(codeNum1, codeNum2 int) string {
	return codeVal[codeNum1] + codeVal[codeNum2]
}

func getValIn() int {
	var val int

	for {
		fmt.Print("Введите исходную валюту 0-USD 1-EUR 2-RUB 3-KGS 4-CNY: ")
		fmt.Scan(&val)

		if val >= 0 && val < COUNT_VALS {
			return val
		} else {
			fmt.Println(ERR_MSG)
		}
	}
}

func getValOut(fromCurrency int) int {
	var val int
	var listCodeVal string

	// убираем из списка код исходной валюты
	for i := 0; i < COUNT_VALS; i++ {
		if i != fromCurrency {
			listCodeVal = fmt.Sprintf("%s %d-%s", listCodeVal, i, getCodeVal(i))
		}
	}

	for {
		fmt.Print("Введите целевую валюту", listCodeVal, ": ")
		fmt.Scan(&val)

		if val < 0 && val >= COUNT_VALS {
			// ошибка ввода
			fmt.Println(ERR_MSG)
			continue
		}

		if val == fromCurrency {
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

func calculateCurs(sumIn float64, fromCurrency, toCurrency int) {
	// заданы курсы базовых пар валют
	mCurs := map[string]float64{"USDEUR": 0.9667, "USDRUB": 101.68, "USDKGS": 87.0, "USDCNY": 7.33}
	var newCurs float64

	if fromCurrency == 0 {
		newCurs = mCurs[getCodeVal2(VAL_USD, toCurrency)]
	} else if toCurrency == 0 {
		newCurs = 1.0 / mCurs[getCodeVal2(VAL_USD, fromCurrency)]
	} else {
		// в исходной или в целевой валюте нет USD
		newCurs = mCurs[getCodeVal2(VAL_USD, toCurrency)] / mCurs[getCodeVal2(VAL_USD, fromCurrency)]

	}

	fmt.Printf("Результат конвертации %.2f %s в %s = %.2f по кросс-курсу %.5f\n",
		sumIn, getCodeVal(fromCurrency), getCodeVal(toCurrency), sumIn*newCurs, newCurs)
}
