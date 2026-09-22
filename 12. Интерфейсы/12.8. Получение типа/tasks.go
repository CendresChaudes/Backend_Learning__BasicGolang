package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("vault"))

	fmt.Println("Задание 2")
	fmt.Println(task2(8))

	fmt.Println("Задание 3")
	fmt.Println(task3("go"))
}

// Задание 1
// Достань из v строку формой с ok: value, ok := v.(string).
// Если ok равен true, верни строку. Иначе верни пустую строку.
func task1(v any) string {
	return ""
}

// Задание 2
// Достань из v число формой с ok: n, ok := v.(int).
// Если ok равен true, верни число. Иначе верни 0.
func task2(v any) int {
	return 0
}

// Задание 3
// Проверь, что внутри v строка. Текст не нужен, на его месте _.
// Верни ok этой проверки.
func task3(v any) bool {
	return false
}
