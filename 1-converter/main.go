package main

import (
	"fmt"
)

func main() {
	const USD_IN_EUR = 0.95
	const USD_IN_RUB = 91.75

	EUR_IN_RUB := USD_IN_RUB / USD_IN_EUR

	fmt.Print(EUR_IN_RUB)
}
