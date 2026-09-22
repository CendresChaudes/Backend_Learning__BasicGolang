package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("clear"))

	fmt.Println("Задание 2")
	fmt.Println(task2(200))

	fmt.Println("Задание 3")
	fmt.Println(task3(12, "C"))
}

// Задание 1
// По короткому коду погоды верни русское слово.
// clear -> ясно, rain -> дождь, snow -> снег.
// Для любого другого кода верни пустую строку.
func task1(kind string) string {
	return ""
}

// Задание 2
// По коду статуса верни короткую метку.
// 200 -> "ok", 404 -> "missing", 500 -> "down".
// Для любого другого code верни "unknown".
func task2(code int) string {
	return ""
}

// Задание 3
// Собери строку через fmt.Sprintf.
// Образец: "%d %s". Сначала temp, потом units.
// Для 12 и C получится "12 C".
func task3(temp int, units string) string {
	return ""
}
