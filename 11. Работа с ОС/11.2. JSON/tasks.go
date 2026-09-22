package main

import "fmt"

type Account struct {
	Login    string `json:"login"`
	Secret   string `json:"secret"`
	Nickname string `json:"nickname,omitempty"`
}

type Entry struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3(`[{"name":"mail","secret":"a"},{"name":"vpn","secret":"b"}]`))
}

// Задание 1
// Собери JSON для Account{Login: "ann", Secret: "qwe"}.
// Nickname оставь пустым.
// Верни текст JSON. Если Marshal вернул ошибку, верни пустую строку.
func task1() string {
	return ""
}

// Задание 2
// Собери JSON для Account{Login: "ann", Nickname: "ani"}.
// Secret оставь пустым.
// Верни текст JSON. Если Marshal вернул ошибку, верни пустую строку.
func task2() string {
	return ""
}

// Задание 3
// Разбери JSON-массив Entry из text.
// Верни Secret записи с Name "vpn".
// Если такой записи нет или Unmarshal вернул ошибку, верни пустую строку.
func task3(text string) string {
	return ""
}
