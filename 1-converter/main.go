package main

import "fmt"

const (
	USDtoEUR = 0.93
	USDtoRUB = 72.10
	EURtoRUB = USDtoRUB / USDtoEUR
)

func main() {
	result := 50 * EURtoRUB
	fmt.Printf("50 EUR равняется %.2f RUB\n", result)
}