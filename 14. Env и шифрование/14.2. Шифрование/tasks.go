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
// Опиши метод Decrypt(encoded string) (string, error), как в теории.
// Тем же ключом расшифруй hex-строку из задания 1.
// Если Decrypt вернул ошибку, верни пустую строку. Иначе верни секрет.
func task2() string {
	return ""
}

// Задание 3
// Вызови Decrypt со строкой "abcd".
// Верни текст ошибки. Для коротких данных это "короткие данные".
func task3() string {
	return ""
}
