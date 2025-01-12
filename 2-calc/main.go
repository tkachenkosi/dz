package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var opers = []string{"AVG", "SUM", "MED", "EXIT"}

func main() {
	fmt.Println("ИССЛЕДОВАНИЕ ТРАНЗАКЦИЙ")

	if oper := getOper(); oper != "EXIT" {
		if tr := getTransact(); len(tr) > 1 {
			fmt.Println("Рсчет", oper, calcValue(oper, tr))
		} else {
			fmt.Println("Для расчета мало данных или не коректно введены транзакции.")
		}
	} else {
		fmt.Println("Выход из программы.")
	}
}

func getOper() string {
	var opr string

	for {
		fmt.Print("Введите операцию ", opers, ": ")
		fmt.Scan(&opr)
		if slices.Contains(opers, opr) {
			return opr
		} else {
			fmt.Println("Ошибка! Введите операцию правильно.")
		}
	}
}

func getTransact() []float64 {
	var values string
	tr := make([]float64, 0, 5)

	fmt.Print("Введите через запятую и без пробелов суммы транзакций: ")
	fmt.Scan(&values)

	numbers := strings.Split(values, ",")

	for _, n := range numbers {
		if v, err := strconv.ParseFloat(n, 64); err == nil {
			tr = append(tr, v)
		}
	}

	return tr
}

func calcValue(oper string, tr []float64) float64 {
	var value float64

	switch oper {
	case opers[0]:
		var count int
		for _, v := range tr {
			if v != 0 {
				value += v
				count++
			}
		}
		value = value / float64(count)
	case opers[1]:
		for _, v := range tr {
			value += v
		}
	case opers[2]:
		length := len(tr)
		slices.Sort(tr)
		if length%2 == 0 {
			value = (tr[length/2] + tr[length/2-1]) / 2
		} else {
			value = tr[length/2]
		}
	}

	return value
}
