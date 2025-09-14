package main

import (
	"fmt"
	"strconv"
	"strings"
)

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.92,
	"RUB": 85.50,
}

func main() {
	fmt.Println("=== Currency Converter ===")

	source := getCurrency("Enter the source currency (USD, EUR, RUB): ", &rates)
	amount := getNumber("Enter the amount: ")
	target := getCurrency("Enter the target currency (USD, EUR, RUB): ", &rates)

	result := calculate(amount, source, target, &rates)

	fmt.Printf("\n%.2f %s = %.2f %s\n", amount, source, result, target)
}

func getCurrency(prompt string, rates *map[string]float64) string {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)
		input = strings.ToUpper(input)

		for k := range *rates {
			if k == input {
				return input
			}
		}
		fmt.Println("Invalid currency. Please choose from", keys(rates))
	}
}

func getNumber(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil || value <= 0 {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		return value
	}
}

func calculate(amount float64, cur1 string, cur2 string, rates *map[string]float64) float64 {
	usd := amount / (*rates)[cur1]
	return usd * (*rates)[cur2]
}

func keys(m *map[string]float64) []string {
	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	return keys
}
