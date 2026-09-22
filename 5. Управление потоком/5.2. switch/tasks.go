package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("GET"))

	fmt.Println("Задание 2")
	fmt.Println(task2(404))

	fmt.Println("Задание 3")
	fmt.Println(task3(6))
}

// Задание 1
// По методу HTTP верни действие.
// GET -> "чтение", POST -> "запись".
// Для любого другого method верни "неизвестно".
func task1(method string) string {
	return ""
}

// Задание 2
// По коду статуса верни метку.
// 200 -> "ok", 404 -> "missing", 500 -> "down".
// Для любого другого code верни "unknown".
func task2(code int) string {
	return ""
}

// Задание 3
// Если day равен 6 или 7, верни "выходной".
// Если day от 1 до 5, верни "работа".
// Для любого другого day верни "нет".
func task3(day int) string {
	return ""
}
