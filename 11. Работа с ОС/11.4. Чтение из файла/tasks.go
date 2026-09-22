package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("note.txt"))

	fmt.Println("Задание 2")
	fmt.Println(task2("note.txt"))

	fmt.Println("Задание 3")
	fmt.Println(task3("note.txt"))
}

// Задание 1
// Прочитай файл path целиком. Верни его содержимое строкой.
// Если прочитать не удалось, верни пустую строку.
func task1(path string) string {
	return ""
}

// Задание 2
// Прочитай файл path. Верни число байтов в нём.
// Если прочитать не удалось, верни 0.
func task2(path string) int {
	return 0
}

// Задание 3
// Если файла нет, верни "нет файла".
// Если файл есть, верни его содержимое строкой.
func task3(path string) string {
	return ""
}
