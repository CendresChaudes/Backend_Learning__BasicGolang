package main

import "fmt"

type Entry struct {
	Name   string
	Secret string
}

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Опиши Encrypter, NewEncrypter и Encrypt, как в теории прошлого урока.
// Ключ прочитай через os.Getenv("VAULT_KEY").
// Зашифруй секрет "b2".
// Если ключа нет или Encrypt вернул ошибку, верни пустую строку.
// Иначе верни hex-строку.
func task1() string {
	return ""
}

// Задание 2
// Ключ снова прочитай из VAULT_KEY.
// Опиши Decrypt и расшифруй эту hex-строку:
// 3031323334353637383961626921247a00f5e1224ea100706d6d7699c2a9
// При ошибке верни пустую строку. Иначе верни секрет.
func task2() string {
	return ""
}

// Задание 3
// Собери Entry{Name: "mail", Secret: "anna"}.
// Ключ прочитай из VAULT_KEY.
// Зашифруй Secret и сразу расшифруй результат.
// Если открытый текст снова равен "anna", верни Name записи.
// Если ключа нет или текст не совпал, верни пустую строку.
func task3() string {
	return ""
}
