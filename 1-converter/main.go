package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Currency = string
type CurrencyRate = float64
type CurrencyRateMap = map[Currency]CurrencyRate
type ExchangeRateMap = map[Currency]CurrencyRateMap
type CurrencySet = map[Currency]bool

var exchangeRate = ExchangeRateMap{
	"USD": {
		"EUR": 0.93,
		"RUB": 72.10,
	},
	"EUR": {
		"USD": 1 / 0.93,
		"RUB": 72.10 / 0.93,
	},
	"RUB": {
		"USD": 1 / 72.10,
		"EUR": 0.93 / 72.10,
	},
}	

func getAmountFromInput() float64 {
	var temp string
	for {
		fmt.Print("Введите количество валюты, которую хотите обменять:")
		fmt.Scan(&temp)
		amount, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			fmt.Println("Введённое значение должно быть числом")
		} else if amount <= 0 {
			fmt.Println("Введённое значение должно быть положительным")
		} else {
			return amount
		}
	}
}

func getCurrencyFromInput(availableCurrencies CurrencySet) string {
	var currency string
	availableCurrenciesStr := strings.Join(func() []string {
		keys := make([]string, 0, len(availableCurrencies))
		for k := range availableCurrencies {
			keys = append(keys, k)
		}
		return keys
	}(), ", ")
	fmt.Println("Доступные валюты:", availableCurrenciesStr)
	for {
		fmt.Print("Введите валюту: ")
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)
		if availableCurrencies[currency] {
			return currency
		} else {
			fmt.Println("Неверное значение для валюты. Возможные варианты: ", availableCurrenciesStr)
		}
	}
}

func getFirstCurrency() string {
	fmt.Println("Какую валюту хотите обменять?")
	return getCurrencyFromInput(func () CurrencySet {
		keys := make(CurrencySet)
		for k := range exchangeRate {
			keys[k] = true
		}
		return keys
	}())
}

func getSecondCurrency(firstCurrency string) string {
	fmt.Println("Какую валюту хотите получить?")
	return getCurrencyFromInput(func () CurrencySet {
		keys := make(CurrencySet)
		for k := range exchangeRate[firstCurrency] {
			if k != firstCurrency {
				keys[k] = true
			}
		}
		return keys
	}())
}


func main() {
	fmt.Println("--- Добро пожаловать в калькулятор валют! ---")
	firstCurrency := getFirstCurrency()
	secondCurrency := getSecondCurrency(firstCurrency)
	amount := getAmountFromInput()
	fmt.Printf(
		"%.2f %s равняется %.2f %s",
		amount, firstCurrency,
		amount*exchangeRate[firstCurrency][secondCurrency], secondCurrency,
	)
}
