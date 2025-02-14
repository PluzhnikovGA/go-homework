package main

import (
	"fmt"
)

func main() {
	fmt.Print("__ Currency conversion __\n")

	amount, yourCurrency, currencyToConvert := inputData()

	result := convertMoney(amount, yourCurrency, currencyToConvert)

	fmt.Printf("You get %.2f %s", result, currencyToConvert)
}

func inputData() (amount float64, yourCurrency string, currencyToConvert string){
	fmt.Print("Enter the amount of money you want to convert: ")
	fmt.Scan(&amount)
	fmt.Print("What currency do you have money in: ")
	fmt.Scan(&yourCurrency)
	fmt.Print("What currency do you want to convert your money into: ")
	fmt.Scan(&currencyToConvert)

	return
}

func convertMoney(amount float64, yourCurrency string, currencyToConvert string) (result float64) {
	const USD_IN_EUR = 0.95
	const USD_IN_RUB = 91.75
	
	return 
}
