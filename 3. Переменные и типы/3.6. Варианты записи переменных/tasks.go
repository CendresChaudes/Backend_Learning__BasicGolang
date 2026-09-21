package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Объяви число 42 тремя способами:
// через var с явным типом, через var без типа и через :=.
// Верни сумму трёх переменных.
func task1() int {
	return 0
}

// Задание 2
// Одной строкой создай language "Go" и year 2009.
// Верни fmt.Sprintf("%s %d", language, year).
func task2() string {
	return ""
}

// Задание 3
// Блоком var ( ) объяви host "localhost", port 8080 и debug false.
// Верни fmt.Sprintf("%s %d %t", host, port, debug).
func task3() string {
	return ""
}
