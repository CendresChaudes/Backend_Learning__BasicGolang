package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1(200))

	fmt.Println("Задание 2")
	fmt.Println(task2(8080))

	fmt.Println("Задание 3")
	fmt.Println(task3("go", "Go"))
}

// Задание 1
// Верни true, если code равен 200.
func task1(code int) bool {
	return false
}

// Задание 2
// Верни true, если port строго больше 1024.
func task2(port int) bool {
	return false
}

// Задание 3
// Верни true, если строки разные.
func task3(left, right string) bool {
	return false
}
