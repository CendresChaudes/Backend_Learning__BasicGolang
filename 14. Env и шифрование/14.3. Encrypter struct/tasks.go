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
// Опиши type Encrypter struct с полем Key string.
// Опиши func NewEncrypter(key string) Encrypter.
// Опиши метод KeySize() int. Он возвращает len(e.Key).
// Верни KeySize для ключа "1234567890123456".
func task1() int {
	return 0
}

// Задание 2
// Опиши метод Mode() string.
// 16 байт — "aes-128", 24 байта — "aes-192", 32 байта — "aes-256".
// Другая длина — пустая строка.
// Верни Mode для ключа "1234567890123456".
func task2() string {
	return ""
}

// Задание 3
// Верни Mode для ключа из 32 символов: "12345678901234567890123456789012".
func task3() string {
	return ""
}
