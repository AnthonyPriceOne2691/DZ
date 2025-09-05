package main

import "fmt"

func main() {

	const usdToEur = 0.92
	const usdToRub = 85.50

	eurToRub := usdToRub / usdToEur

	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)
}

func getUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func calculate() float64 {
	return 0
}