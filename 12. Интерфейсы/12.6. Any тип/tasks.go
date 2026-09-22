package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("vault"))

	fmt.Println("Задание 2")
	fmt.Println(task2(120))

	fmt.Println("Задание 3")
	fmt.Println(task3("vault"))
}

// Задание 1
// Верни текст значения v через fmt.Sprint.
func task1(v any) string {
	return ""
}

// Задание 2
// Верни длину текста, который fmt.Sprint собирает из v.
func task2(v any) int {
	return 0
}

// Задание 3
// Верни true, если в v лежит значение, а не nil.
func task3(v any) bool {
	return false
}
