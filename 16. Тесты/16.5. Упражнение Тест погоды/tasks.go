package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("Moscow", 12))

	fmt.Println("Задание 2")
	fmt.Println(task2("Moscow", "C"))

	fmt.Println("Задание 3")
	fmt.Println(task3("Moscow"))
}

// Задание 1
// Собери отчёт через fmt.Sprintf.
// Образец: "%s: %d C". Сначала город, потом температура.
// Для Moscow и 12 получится "Moscow: 12 C".
func task1(city string, temp int) string {
	return ""
}

// Задание 2
// Собери url.Values.
// Поставь city и units через Set.
// Верни результат Encode.
// Для Moscow и C получится "city=Moscow&units=C".
func task2(city, units string) string {
	return ""
}

// Задание 3
// Собери url.Values с city и units "C".
// Верни адрес "https://api.weather.local/v1?", затем Encode.
// Для Moscow получится "https://api.weather.local/v1?city=Moscow&units=C".
func task3(city string) string {
	return ""
}
