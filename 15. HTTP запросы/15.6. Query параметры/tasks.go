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
// Собери url.Values.
// Поставь city = Moscow и units = C через Set.
// Верни результат Encode.
func task1() string {
	return ""
}

// Задание 2
// Собери те же пары: city = Moscow, units = C.
// Верни полный адрес: "https://api.weather.local/v1", затем "?", затем Encode.
func task2() string {
	return ""
}

// Задание 3
// Разбери строку "city=Kazan&units=F" через url.ParseQuery.
// Если ParseQuery вернул ошибку, верни пустую строку.
// Иначе верни значение ключа city.
func task3() string {
	return ""
}
