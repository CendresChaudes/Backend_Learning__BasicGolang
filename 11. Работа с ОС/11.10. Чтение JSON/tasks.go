package main

import "fmt"

type Entry struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("vault.json"))

	fmt.Println("Задание 2")
	fmt.Println(task2("vault.json"))

	fmt.Println("Задание 3")
	fmt.Println(task3("vault.json"))
}

// Задание 1
// Прочитай JSON-массив из path в срез Entry.
// Верни число записей.
// Если прочитать или разобрать не удалось, верни 0.
func task1(path string) int {
	return 0
}

// Задание 2
// Прочитай JSON-массив из path.
// Верни Secret записи с Name "vpn".
// Если такой записи нет или файл не разобран, верни пустую строку.
func task2(path string) string {
	return ""
}

// Задание 3
// Если файла нет, верни "нет файла".
// Если текст не JSON, верни "битый json".
// Если массив разобран, верни Name первой записи.
func task3(path string) string {
	return ""
}
