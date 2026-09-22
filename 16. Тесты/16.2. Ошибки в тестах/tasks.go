package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(""))

	fmt.Println("Задание 2")
	fmt.Println(task2(404))

	fmt.Println("Задание 3")
	fmt.Println(task3(""))
}

// Задание 1
// Если city — пустая строка, создай ошибку errors.New("пустой город")
// и верни текст метода Error.
// Если city не пустой, ошибки нет: верни пустую строку.
func task1(city string) string {
	return ""
}

// Задание 2
// Если code не равен 200, создай ошибку errors.New("плохой статус")
// и верни текст метода Error.
// Если code равен 200, верни "ok".
func task2(code int) string {
	return ""
}

// Задание 3
// Если body — пустая строка, создай ошибку errors.New("пустое тело")
// и верни текст метода Error.
// Если body не пустой, верни сам body.
func task3(body string) string {
	return ""
}
