package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(200))

	fmt.Println("Задание 2")
	fmt.Println(task2(20))

	fmt.Println("Задание 3")
	fmt.Println(task3(4))
}

// Задание 1
// Если code равен 200, верни "ok". Иначе верни "ошибка".
func task1(code int) string {
	return ""
}

// Задание 2
// Если age больше или равен 18, верни true. Иначе верни false.
func task2(age int) bool {
	return false
}

// Задание 3
// Если n больше 0, верни n, умноженное на 2. Иначе верни 0.
func task3(n int) int {
	return 0
}
