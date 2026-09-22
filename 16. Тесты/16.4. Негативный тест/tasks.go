package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("Moscow"))

	fmt.Println("Задание 2")
	fmt.Println(task2(404))

	fmt.Println("Задание 3")
	fmt.Println(task3("Moscow", "ясно"))
}

// Задание 1
// Если city — пустая строка, верни пустую строку.
// Иначе верни "city=" и следом сам city.
// Для Moscow получится "city=Moscow".
func task1(city string) string {
	return ""
}

// Задание 2
// Верни true, только если code равен 200.
// Для любого другого code верни false.
func task2(code int) bool {
	return false
}

// Задание 3
// Если city или text — пустая строка, верни "нет данных".
// Иначе верни отчёт: город, ": ", текст.
// Для Moscow и ясно получится "Moscow: ясно".
func task3(city, text string) string {
	return ""
}
