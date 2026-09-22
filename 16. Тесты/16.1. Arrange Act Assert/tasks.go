package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("Moscow", "ясно"))

	fmt.Println("Задание 2")
	fmt.Println(task2(200))

	fmt.Println("Задание 3")
	fmt.Println(task3("Moscow: ясно", "Moscow: ясно"))
}

// Задание 1
// Собери строку отчёта: город, затем ": ", затем текст погоды.
// Для Moscow и ясно получится "Moscow: ясно".
func task1(city, text string) string {
	return ""
}

// Задание 2
// Это значение, которое тест потом положит в want.
// Если code равен 200, верни "ok".
// Для любого другого code верни "fail".
func task2(code int) string {
	return ""
}

// Задание 3
// Это шаг Assert: сравни две строки.
// Верни true, если got и want равны.
// Иначе верни false.
func task3(got, want string) bool {
	return false
}
