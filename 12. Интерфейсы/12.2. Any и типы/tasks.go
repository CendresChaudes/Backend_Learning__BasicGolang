package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("vault"))

	fmt.Println("Задание 2")
	fmt.Println(task2("vault"))

	fmt.Println("Задание 3")
	fmt.Println(task3(7))
}

// Задание 1
// По типу v верни имя.
// string -> "string", int -> "int".
// Для любого другого типа верни "other".
func task1(v any) string {
	return ""
}

// Задание 2
// Если v — string, верни эту строку.
// Если тип другой, верни пустую строку.
func task2(v any) string {
	return ""
}

// Задание 3
// Если v — int, верни это число.
// Если тип другой, верни -1.
func task3(v any) int {
	return 0
}
