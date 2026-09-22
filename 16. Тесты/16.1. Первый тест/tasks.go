package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("Moscow", "ясно"))

	fmt.Println("Задание 2")
	fmt.Println(task2(200))

	fmt.Println("Задание 3")
	fmt.Println(task3(4))
}

// Задание 1
// Верни город, двоеточие, пробел и text.
// Для Moscow и "ясно" получится "Moscow: ясно".
func task1(city, text string) string {
	return ""
}

// Задание 2
// Если code равен 200, верни "ok". Иначе верни "ошибка".
func task2(code int) string {
	return ""
}

// Задание 3
// Верни n, умноженное на 2.
func task3(n int) int {
	return 0
}
