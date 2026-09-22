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
// Адрес: https://api.weather.local/v1?city=Moscow
// Статус: 200.
// Верни debug-строку через fmt.Sprintf.
// Формат: "GET %s -> %d".
func task1() string {
	return ""
}

// Задание 2
// Тот же адрес. Статус: 404.
// Если статус не равен 200, верни debug-строку "GET %s -> %d".
// Если статус равен 200, верни отчёт "Moscow: ясно".
func task2() string {
	return ""
}

// Задание 3
// Тот же адрес. Статус: 200. Текст погоды: "ясно".
// Если статус не равен 200, верни debug-строку.
// Если статус равен 200, верни отчёт "Moscow: ясно".
func task3() string {
	return ""
}
