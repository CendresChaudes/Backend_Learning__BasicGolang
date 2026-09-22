package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(map[string]int{"tea": 10, "coffee": 20}))

	fmt.Println("Задание 2")
	fmt.Println(task2(map[string]int{"a": 5, "b": 11, "c": 40}))

	fmt.Println("Задание 3")
	fmt.Println(task3(map[string]int{"ok": 200, "missing": 404}))
}

// Задание 1
// Верни сумму всех значений в prices. Обойди map через range.
func task1(prices map[string]int) int {
	return 0
}

// Задание 2
// Верни, сколько значений в items строго больше 10. Обойди map через range.
func task2(items map[string]int) int {
	return 0
}

// Задание 3
// Верни ключ, у которого значение равно 404.
// Если такого значения нет, верни пустую строку.
func task3(codes map[string]int) string {
	return ""
}
