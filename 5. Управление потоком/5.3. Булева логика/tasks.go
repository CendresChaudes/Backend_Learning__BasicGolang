package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(true, true))

	fmt.Println("Задание 2")
	fmt.Println(task2(201))

	fmt.Println("Задание 3")
	fmt.Println(task3(false))
}

// Задание 1
// Верни true, только если admin и active оба true.
func task1(admin bool, active bool) bool {
	return false
}

// Задание 2
// Верни true, если code равен 200 или 201.
func task2(code int) bool {
	return false
}

// Задание 3
// Верни противоположное значение: true, если blocked равен false.
func task3(blocked bool) bool {
	return false
}
