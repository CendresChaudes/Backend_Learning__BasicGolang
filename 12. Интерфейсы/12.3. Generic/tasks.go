package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("vault"))

	fmt.Println("Задание 2")
	fmt.Println(task2(7, 7))

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Опиши func Echo[T any](v T) T. Она возвращает v.
// task1 вызывает Echo и возвращает результат.
func task1(v string) string {
	return ""
}

// Задание 2
// Опиши func Same[T comparable](a, b T) bool.
// Она возвращает true, когда a равен b.
// task2 вызывает Same и возвращает результат.
func task2(a, b int) bool {
	return false
}

// Задание 3
// Опиши type Box[T any] struct с полем Value T.
// Создай Box[int] со значением 7 и верни это поле.
func task3() int {
	return 0
}
