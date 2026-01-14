package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
Нужно создать приложение, которое умеет делать простые расчёты над числами.
Приложение должно:
- Принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
- Принимает неограниченное число чисел через запятую (2, 10, 9)
- Разбивает строку чисел по запятым и затем делает расчёт в зависимости от операции выводя результат
*/

func calculateAverage(numbers []int) float64 {
	// Calculate average of a slice of integers
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}

func calculateSum(numbers []int) int {
	// Calculate sum of a slice of integers
	var sum int
	for _, number := range numbers {
		sum += number
	}	
	return sum
}

func calculateMedian(numbers []int) float64 {
	// Calculate median of a slice of integers
	n := len(numbers)
	sortedNumbers := make([]int, n)
	copy(sortedNumbers, numbers)
	sort.Ints(sortedNumbers)
	if n%2 == 0 {
		return float64(sortedNumbers[n/2-1]+sortedNumbers[n/2]) / 2.0
	}
	return float64(sortedNumbers[n/2])
}

func calculate(operation string, numbers []int) (result float64) {
	// Perform calculation based on the operation
	switch operation {
	case "AVG":
		result = calculateAverage(numbers)
	case "SUM":
		result = float64(calculateSum(numbers))
	case "MED":
		result = calculateMedian(numbers)
	}
	return 
}

func getOperationFromInput() (operation string) {
	// Get operation from user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите операцию (AVG, SUM, MED): ")
		scanner.Scan()
		operation = scanner.Text()
		operation = strings.ToUpper(strings.TrimSpace(operation))
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			return operation
		} else {
			fmt.Println("Неверная операция. Попробуйте ещё раз.")
		}
	}
}

func getNumbersFromInput() (numbers []int) {
	// Get numbers from user input
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	out:
	for {
		fmt.Print("Введите числа через запятую: ")
		scanner.Scan()
		input = scanner.Text()
		numbers = []int{}
		for numStr := range strings.SplitSeq(input, ",") {
			var num int
			_, err := fmt.Sscanf(strings.TrimSpace(numStr), "%d", &num)
			if err != nil {
				fmt.Printf("Ошибка при парсинге числа: \"%s\"\n", numStr)
				continue out
			}
			numbers = append(numbers, num)
		}
		if len(numbers) > 0 {
			return numbers
		} else {
			fmt.Println("Необходимо ввести хотя бы одно число.")
		}
	}
}

func main() {
	// Main function to run the calculator application
	operation := getOperationFromInput()
	numbers := getNumbersFromInput()
	result := calculate(operation, numbers)
	fmt.Printf("Результат операции %s: %.2f\n", operation, result)
}