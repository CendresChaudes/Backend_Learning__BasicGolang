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
// Опиши Note с полями Title string и Pages int, в таком порядке.
// Собери JSON для Title "Go" и Pages 3 через json.Marshal.
// Верни этот текст.
func task1() string {
	return ""
}

// Задание 2
// Используй тип Note из задания 1.
// Разбери JSON {"Title":"Go","Pages":3} через json.Unmarshal.
// Верни Pages.
func task2() int {
	return 0
}

// Задание 3
// Разбери текст {"Title": через json.Unmarshal.
// Верни true, если вызов вернул ошибку.
func task3() bool {
	return false
}
