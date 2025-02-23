package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var result float64

	method := getMethod(reader)

	numbers := getNumbersList(reader)

	switch {
	case method == "AVG":
		result = averageAllNumbers(numbers)
	case method == "SUM":
		result = summaAllNumbers(numbers)
	default:
		result = medianAllNumbers(numbers)
	}

	fmt.Printf("You chose '%s' and got the result: %.5f", method, result)
}

func getMethod(reader *bufio.Reader) (method string) {
	methods := []string{"AVG", "SUM", "MED"}
	strMethods := strings.Join(methods, ", ")

	fmt.Printf("Enter a method from this list [%s]: ", strMethods)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input, please try again.")
			continue
		}
		method = strings.ToUpper(strings.TrimSpace(input))

		if slices.Index(methods, method) >= 0 {
			return method
		}

		fmt.Printf("You entered an incorrect method, you need to choice a method from this list [%s]: ", strMethods)
	}
}

func getNumbersList(reader *bufio.Reader) (numbers []float64) {
	fmt.Print("Enter a list of numbers separated by commas: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input. Please try again.")
		return getNumbersList(reader)
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	for _, number := range parts {
		num, err := strconv.ParseFloat(strings.TrimSpace(strings.Trim(number, ",")), 64)
		if err != nil {
			fmt.Printf("Error: '%s' is not a number. Repeat enter. \n", number)
			return getNumbersList(reader)
		}

		numbers = append(numbers, num)
	}

	return numbers
}

func summaAllNumbers(numbers []float64) (result float64) {
	for _, value := range numbers {
		result += value
	}

	return result
}

func averageAllNumbers(numbers []float64) (result float64) {
	return summaAllNumbers(numbers) / float64(len(numbers))
}

func medianAllNumbers(numbers []float64) (result float64) {
	length := len(numbers)

	if length == 0 {
		return 0
	}

	sort.Float64s(numbers)

	mid := length / 2

	if (length % 2 == 1) {
		return numbers[mid]
	}

	return (numbers[mid - 1] + numbers[mid]) /2
}