package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	USDtoEUR = 0.93
	USDtoRUB = 72.10
	EURtoRUB = USDtoRUB / USDtoEUR
	EURtoUSD = 1 / USDtoEUR
	RUBtoUSD = 1 / USDtoRUB
	RUBtoEUR = USDtoEUR / USDtoRUB
)

func getAmountFromInput() (amount float64) {
	var temp string
	var err error
	for  {
		fmt.Print("Введите количество валюты, которую хотите обменять:")
		fmt.Scan(&temp)
		amount, err = strconv.ParseFloat(temp, 64)
		if err != nil {
			fmt.Println("Введённое значение должно быть числом")
		} else if amount <= 0 {
			fmt.Println("Введённое значение должно быть положительным")
		} else {
			break
		}
	}
	return
}

func getCurrencyFromInput() (currency string) {
	for {
		fmt.Print("Введите валюту: ")
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)
		if currency == "RUB" || currency == "EUR" || currency == "USD" {
			break
		} else {
			fmt.Print("Неверное значение для валюты или такая валюта не поддерживается")
		}
	}
	return
}

func getFirstCurrency() (firstCurrency string) {
	fmt.Println("Какую валюту хотите обменять?\nВозможные варианты: RUB, USD, EUR")
	firstCurrency = getCurrencyFromInput()
	return 
}

func getSecondCurrency(firstCurrency string) (secondCurrency string) {
	var availableCurrencies string
	switch firstCurrency {
	case "RUB":
		availableCurrencies = "USD, EUR"
	case "USD":
		availableCurrencies = "RUB, EUR"
	default:
		availableCurrencies = "RUB, USD"
	}
	fmt.Println("Какую валюту хотите получить?\nВозможные варианты:", availableCurrencies)
	for secondCurrency = getCurrencyFromInput(); secondCurrency == firstCurrency; secondCurrency = getCurrencyFromInput() {
		fmt.Println("Валюты не могут совпадать\nВозможные варианты:", availableCurrencies)
	}
	return
}


func convert(amount float64, firstCurrency, secondCurrency string) (result float64) {
	switch firstCurrency {
	case "RUB":
		switch secondCurrency {
		case "USD":
			result = amount * RUBtoUSD
		case "EUR":
			result = amount * RUBtoEUR
		}
	case "USD":
		switch secondCurrency {
		case "RUB":
			result = amount * USDtoRUB
		case "EUR":
			result = amount * USDtoEUR
		}
	case "EUR":
		switch secondCurrency {
		case "RUB":
			result = amount * EURtoRUB
		case "USD":
			result = amount * EURtoUSD
		}
	}
	return
}

func main() {
	fmt.Println("--- Добро пожаловать в калькулятор валют! ---")
	firstCurrency := getFirstCurrency()
	secondCurrency := getSecondCurrency(firstCurrency)
	amount := getAmountFromInput()
	fmt.Printf(
		"%.2f %s равняется %.2f %s", 
		amount, firstCurrency,
		convert(amount, firstCurrency, secondCurrency), secondCurrency,
	)
}