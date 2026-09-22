package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("account.json"))

	fmt.Println("Задание 2")
	fmt.Println(task2("account.json", "ann"))

	fmt.Println("Задание 3")
	fmt.Println(task3("account.json"))
}

// Задание 1
// Опиши Account: Login с тегом json "login", Secret с тегом json "secret".
// Запиши в path JSON для Login "ann" и Secret "qwe". Права 0644.
// Верни этот JSON текстом.
func task1(path string) string {
	return ""
}

// Задание 2
// Опиши отдельный тип LoginFile с полем Login и тегом json "login".
// Если login пустой, верни "пустой логин" и файл не создавай.
// Иначе запиши в path JSON этого объекта. Права 0644.
// При успехе верни пустую строку.
func task2(path, login string) string {
	return ""
}

// Задание 3
// Опиши отдельный тип Service: Login с тегом json "login", Port с тегом json "port".
// Запиши в path JSON для Login "ann" и Port 80. Права 0644.
// Верни число байтов этого JSON.
func task3(path string) int {
	return 0
}
