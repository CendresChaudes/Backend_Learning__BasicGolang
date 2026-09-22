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
// На уровне пакета опиши func Echo[T any](v T) T.
// Она возвращает v.
// В task1 верни Echo("vault").
func task1() string {
	return ""
}

// Задание 2
// Опиши func Count[T any](items []T) int.
// Она возвращает len(items).
// В task2 верни Count от среза строк "a", "b", "c".
func task2() int {
	return 0
}

// Задание 3
// Опиши func HasItem[T any](items []T) bool.
// Она возвращает true, если len(items) больше нуля.
// В task3 верни HasItem от среза int с одним числом 7.
func task3() bool {
	return false
}
