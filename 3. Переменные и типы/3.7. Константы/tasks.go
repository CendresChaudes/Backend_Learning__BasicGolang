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
// Объяви константу appName со значением "wallet".
// Верни её.
func task1() string {
	const appName = "wallet"
	return appName
}

// Задание 2
// Объяви константу maxAttempts со значением 3.
// Объяви переменную attempts со значением 1 и увеличь её на 1.
// Верни attempts + maxAttempts. Константу не переназначай.
func task2() int {
	const maxAttempts = 3
	attempts := 1
	attempts = attempts + 1
	return attempts + maxAttempts
}

// Задание 3
// Блоком const объяви monday, tuesday и wednesday через iota, чтобы номера шли 0, 1, 2.
// Верни значение wednesday.
func task3() int {
	const (
		monday = iota
		tuesday
		wednesday
	)

	return wednesday
}
