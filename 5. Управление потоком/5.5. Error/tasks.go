package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(""))

	fmt.Println("Задание 2")
	fmt.Println(task2(0))

	fmt.Println("Задание 3")
	fmt.Println(task3(-1))
}

// Задание 1
// Если port — пустая строка, создай ошибку errors.New("пустой порт")
// и верни текст метода Error.
// Если port не пустой, ошибки нет: верни пустую строку.
func task1(port string) string {
	return ""
}

// Задание 2
// Если b равен 0, создай ошибку errors.New("деление на ноль")
// и верни текст метода Error.
// Если b не равен 0, верни пустую строку.
func task2(b int) string {
	return ""
}

// Задание 3
// Если n меньше 0, верни текст ошибки errors.New("отрицательное").
// Иначе верни "ok".
func task3(n int) string {
	return ""
}
