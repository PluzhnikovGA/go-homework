package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	currencyList := []string{"USD", "RUB", "EUR"}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("__ Currency conversion __\n")

	yourCurrency, amount, currencyToConvert := inputData(currencyList, reader)

	result := convertMoney(amount, yourCurrency, currencyToConvert)

	fmt.Printf("You'll get %.2f %s", result, currencyToConvert)
}

var currencyRates = map[string]map[string]float64{
	"USD": {"EUR": 0.95, "RUB": 91.07},
	"EUR": {"USD": 1.05, "RUB": 95.61},
	"RUB": {"USD": 0.011, "EUR": 0.01},
}

func inputData(currencyList []string, reader *bufio.Reader) (yourCurrency string, amount float64, currencyToConvert string){
	yourCurrency, currencyList = choiceCurrent(currencyList, "What currency do you have money in", reader)

	amount = getAmount()

	currencyToConvert, _ = choiceCurrent(currencyList, "What currency do you want to convert your money into", reader)

	return yourCurrency, amount, currencyToConvert
}

func convertMoney(amount float64, yourCurrency string, currencyToConvert string) (result float64) {
	if rates, ok := currencyRates[yourCurrency]; ok {
			if rate, ok := rates[currencyToConvert]; ok {
					return amount * rate
			}
	}
	return amount
}

func choiceCurrent(currencyList []string, text string, reader *bufio.Reader) (currency string, newCurrencyList []string) {
	currencyListStr := strings.Join(currencyList, ", ")
	fmt.Printf("%s [%s]: ", text, currencyListStr)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.ToUpper(strings.TrimSpace(input))

		currencyIndex := slices.Index(currencyList, input)

		if currencyIndex >= 0 {
			currencyList = append(currencyList[:currencyIndex], currencyList[currencyIndex+1:]...)
			return input, currencyList
		}

		fmt.Printf("You need to choice currency from this list [%s]: ", currencyListStr)
	}
}

func getAmount() (amount float64) {
	fmt.Print("Enter the amount of money you want to convert: ")
	for {
		_, err := fmt.Scan(&amount)
		if amount > 0 && err == nil {
			return amount
		}
		fmt.Print("You entered an incorrect amount, please repeat: ")
	}
}
