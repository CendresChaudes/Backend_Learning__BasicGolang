package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("10 20 30"))

	fmt.Println("Задание 2")
	fmt.Println(task2("4 5 6 7"))

	fmt.Println("Задание 3")
	fmt.Println(task3("1 4 15 2"))
}

// Задание 1
// В text числа через пробел. Прочитай их по одному циклом и верни сумму.
// Подключи пакет strings. Источник — strings.NewReader, чтение — fmt.Fscan.
func task1(text string) int {
	return 0
}

// Задание 2
// Верни, сколько чисел лежит в text. Читай их тем же циклом.
func task2(text string) int {
	return 0
}

// Задание 3
// Верни первое число, которое строго больше 10.
// Когда нашёл его, выйди через break. Если такого числа нет, верни 0.
func task3(text string) int {
	return 0
}
