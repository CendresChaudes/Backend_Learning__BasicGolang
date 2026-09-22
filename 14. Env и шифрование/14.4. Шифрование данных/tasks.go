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
// Опиши type Encrypter struct с полем Key string и func NewEncrypter(key string) Encrypter.
// Опиши метод Encrypt(plain string) (string, error), как в теории.
// Nonce бери такой: []byte("0123456789ab").
// Зашифруй "vault" ключом "1234567890123456".
// Если Encrypt вернул ошибку, верни пустую строку. Иначе верни hex-строку.
func task1() string {
	return ""
}

// Задание 2
// Тем же ключом и тем же nonce зашифруй "mail".
// При ошибке верни пустую строку.
func task2() string {
	return ""
}

// Задание 3
// Тем же ключом и тем же nonce зашифруй "b2".
// При ошибке верни пустую строку.
func task3() string {
	return ""
}
