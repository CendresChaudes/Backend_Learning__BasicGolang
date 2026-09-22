package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(""))

	fmt.Println("Задание 2")
	fmt.Println(task2("https://api.weather.local", "/v1"))

	fmt.Println("Задание 3")
	fmt.Println(task3(12, "C"))
}

// Задание 1
// Если units — пустая строка, верни "C".
// Иначе верни units как есть.
func task1(units string) string {
	return ""
}

// Задание 2
// Склей базовый адрес: host, затем path.
// Для host "https://api.weather.local" и path "/v1"
// получится "https://api.weather.local/v1".
func task2(host, path string) string {
	return ""
}

// Задание 3
// Собери строку температуры через fmt.Sprintf.
// Образец: "%d %s". Сначала число, потом units.
// Для 12 и C получится "12 C".
func task3(temp int, units string) string {
	return ""
}
