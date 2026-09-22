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
// Создай переменную message через fmt.Sprintf.
// Образец: "пользователь %s", значение: "Анна".
// Верни message.
func task1() string {
	return fmt.Sprintf("пользователь %s", "Анна")
}

// Задание 2
// Через fmt.Sprintf собери "сумма 19.90".
// Глагол %.2f, число 19.9.
func task2() string {
	return fmt.Sprintf("сумма %.2f", 19.9)
}

// Задание 3
// Через fmt.Sprintf и глагол %q получи строку Go в двойных кавычках.
func task3() string {
	return fmt.Sprintf("%q", "Go")
}
