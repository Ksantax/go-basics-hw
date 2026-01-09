package main

import "fmt"

const (
	USDtoEUR = 0.93
	USDtoRUB = 72.10
)

func main() {
	result := (50 / USDtoEUR) * USDtoRUB
	fmt.Printf("50 EUR равняется %.2f RUB\n", result)
}