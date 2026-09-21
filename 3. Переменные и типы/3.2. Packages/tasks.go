package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2("backend"))

	fmt.Println("Задание 3")
	fmt.Println(task3("=", 20))
}

// Задание 1
// Верни имя пакета, в котором лежит этот файл.
// Это программа, а не библиотека.
func task1() string {
	return "main"
}

// Задание 2
// text — строка, которую передают при вызове. Подключать своё значение не нужно.
// Подключи пакет strings и верни text заглавными буквами через strings.ToUpper.
func task2(text string) string {
	return strings.ToUpper(text)
}

// Задание 3
// char — кусок текста, n — сколько раз его повторить.
// Верни результат strings.Repeat.
func task3(char string, n int) string {
	return strings.Repeat(char, n)
}
