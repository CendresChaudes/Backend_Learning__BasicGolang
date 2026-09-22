package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("vault"))

	fmt.Println("Задание 2")
	fmt.Println(task2("vault"))

	fmt.Println("Задание 3")
	fmt.Println(task3("go"))
}

// Задание 1
// Через type switch верни имя типа внутри v.
// Для string верни "string", для int верни "int".
// Для любого другого типа верни "other".
func task1(v any) string {
	return ""
}

// Задание 2
// Через type switch с переменной ветки верни размер.
// Для string верни len. Для int верни само число.
// Для любого другого типа верни 0.
func task2(v any) int {
	return 0
}

// Задание 3
// Через type switch верни true, если внутри v строка.
// Для любого другого типа верни false.
func task3(v any) bool {
	return false
}
