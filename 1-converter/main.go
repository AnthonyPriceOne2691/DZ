package main

import (
	"fmt"
	"strconv"
	"strings"
)

const usdToEur = 0.92
const usdToRub = 85.50

func main() {
	fmt.Println("=== Currency Converter ===")

	source := getCurrency("Enter the source currency (USD, EUR, RUB): ")
	amount := getNumber("Enter the amount: ")
	target := getCurrency("Enter the target currency (USD, EUR, RUB): ")

	result := calculate(amount, source, target)

	fmt.Printf("\n%.2f %s = %.2f %s\n", amount, source, result, target)
}

func getCurrency(prompt string) string {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scan(&input)
		input = strings.ToUpper(input)

		if input == "USD" || input == "EUR" || input == "RUB" {
			return input
		}
		fmt.Println("Invalid currency. Please choose from USD, EUR, RUB.")
	}
}

func getNumber(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid number. Please enter a positive numeric value.")
			continue
		}

		if value <= 0 {
			fmt.Println("Invalid number. Please enter a positive numeric value.")
			continue
		}

		return value
	}
}

func calculate(amount float64, cur1 string, cur2 string) float64 {
	if cur1 == cur2 {
		return amount
	}

	var amountInUsd float64
	switch cur1 {
	case "USD":
		amountInUsd = amount
	case "EUR":
		amountInUsd = amount / usdToEur
	case "RUB":
		amountInUsd = amount / usdToRub
	}

	switch cur2 {
	case "USD":
		return amountInUsd
	case "EUR":
		return amountInUsd * usdToEur
	case "RUB":
		return amountInUsd * usdToRub
	default:
		return 0
	}
}
