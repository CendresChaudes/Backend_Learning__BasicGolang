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
// Опиши type Encrypter, NewEncrypter и метод Decrypt(encoded string) (string, error).
// Расшифруй ключом "1234567890123456" эту hex-строку:
// 3031323334353637383961627d720c0df0d0d984191548689c68bd324d0c89e005
// Если Decrypt вернул ошибку, верни пустую строку. Иначе верни секрет.
func task1() string {
	return ""
}

// Задание 2
// Тем же ключом расшифруй:
// 3031323334353637383961626672100d8bc56da5460662a58605ff7169b89b52
// При ошибке верни пустую строку.
func task2() string {
	return ""
}

// Задание 3
// Тем же ключом расшифруй короткую строку:
// 30313233343536373839616200
// Если Decrypt вернул ошибку, верни "ошибка".
// Если ошибки нет, верни секрет.
func task3() string {
	return ""
}
