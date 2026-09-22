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
// Опиши func Same[T comparable](a, b T) bool.
// Она возвращает a == b.
// В task1 верни Same("go", "go").
func task1() bool {
	return false
}

// Задание 2
// Опиши интерфейс Number с объединением int | int64.
// Опиши func Add[T Number](a, b T) T. Она возвращает a + b.
// В task2 верни Add(2, 5).
func task2() int {
	return 0
}

// Задание 3
// Опиши интерфейс Ordered с объединением int | string.
// Опиши func Larger[T Ordered](a, b T) T.
// Если a > b, верни a. Иначе верни b.
// В task3 верни Larger("disk", "cloud").
func task3() string {
	return ""
}
