package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var operations = map[string]func([]float64) float64{
	"SUM": sum,
	"AVG": avg,
	"MED": median,
}

func main() {
	operation := strings.ToUpper(readInput("Введите операцию (SUM, AVG, MED): "))
	input := readInput("Введите числа (через запятую или пробел): ")

	numbers, err := parseNumbers(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(numbers) == 0 {
		fmt.Println("Нет чисел для расчёта")
		return
	}

	result := operations[operation](numbers)
	fmt.Printf("%s = %.2f\n", operation, result)
}

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func parseNumbers(input string) ([]float64, error) {
	input = strings.ReplaceAll(input, ",", " ")
	parts := strings.Fields(input)

	var numbers []float64
	for _, p := range parts {
		num, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err != nil {
			return nil, fmt.Errorf("некорректное число: %s", p)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func sum(numbers []float64) float64 {
	var total float64
	for _, n := range numbers {
		total += n
	}
	return total
}

func avg(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

func median(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 0 {
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}
