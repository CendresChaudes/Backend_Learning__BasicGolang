package main

import "fmt"

var service = "api"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3(8080))
}

// Задание 1
// service объявлена снаружи функции, на уровне пакета.
// Верни её.
func task1() string {
	return service
}

// Задание 2
// Внутри функции создай новую переменную service со значением "worker".
// Она закрывает внешнюю service только здесь. Верни внутреннюю.
func task2() string {
	service := "worker"
	return service
}

// Задание 3
// port виден только внутри этой функции. Верни его.
func task3(port int) int {
	return port
}
