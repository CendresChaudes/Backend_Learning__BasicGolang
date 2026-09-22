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
// Создай map[string]int с тремя парами: "ok" → 200, "created" → 201, "missing" → 404.
// Верни len этой map.
func task1() int {
	return 0
}

// Задание 2
// Создай map[string]string с парой "go" → "golang".
// Верни значение ключа "go".
func task2() string {
	return ""
}

// Задание 3
// Создай map[string]int с парой "pen" → 15.
// Прочитай ключ "book": его в map нет, поэтому чтение даст 0.
// Прибавь к этому результату 7 и верни сумму.
func task3() int {
	return 0
}
