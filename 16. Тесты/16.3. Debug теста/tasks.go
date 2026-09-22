package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("ясно", "дождь"))

	fmt.Println("Задание 2")
	fmt.Println(task2("TestTask1", true))

	fmt.Println("Задание 3")
	fmt.Println(task3(404, "https://api.weather.local/v1?city=Moscow"))
}

// Задание 1
// Собери текст провала через fmt.Sprintf.
// Образец: получилось %q, нужно %q.
// Сначала подставь got, потом want.
// Для ясно и дождь получится: получилось "ясно", нужно "дождь".
func task1(got, want string) string {
	return ""
}

// Задание 2
// Если ok равно true, верни "PASS", пробел и name.
// Если ok равно false, верни "FAIL", пробел и name.
// Для TestTask1 и true получится "PASS TestTask1".
func task2(name string, ok bool) string {
	return ""
}

// Задание 3
// Если code равен 200, верни "ok".
// Иначе собери debug-строку через fmt.Sprintf.
// Образец: "GET %s -> %d". Сначала rawURL, потом code.
// Для 404 и адреса https://api.weather.local/v1?city=Moscow
// получится "GET https://api.weather.local/v1?city=Moscow -> 404".
func task3(code int, rawURL string) string {
	return ""
}
