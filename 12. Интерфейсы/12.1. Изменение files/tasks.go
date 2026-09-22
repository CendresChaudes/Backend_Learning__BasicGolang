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
// На уровне пакета опиши type FileBox struct с полем Name string.
// Метод Title с получателем-значением FileBox возвращает Name.
// Создай FileBox с именем "vault.json" и верни Title.
func task1() string {
	return ""
}

// Задание 2
// У FileBox метод Mark возвращает строку "file " и сразу за ней Name.
// Создай FileBox с именем "vault.json" и верни Mark.
func task2() string {
	return ""
}

// Задание 3
// У FileBox метод Ready возвращает true, если Name не пустой.
// Создай FileBox с именем "vault.json" и верни Ready.
func task3() bool {
	return false
}
