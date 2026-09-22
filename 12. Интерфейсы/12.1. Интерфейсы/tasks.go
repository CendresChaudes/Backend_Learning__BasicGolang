package main

import "fmt"

type Labeler interface {
	Label() string
}

type Disk struct {
	Name string
}

type Cloud struct {
	Name string
}

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Опиши метод Label() string у Disk. Он возвращает "disk:" и имя.
// Опиши func place(item Labeler) string. Она возвращает item.Label().
// Верни place для Disk с именем "a".
func task1() string {
	return ""
}

// Задание 2
// Опиши метод Label() string у Cloud. Он возвращает "cloud:" и имя.
// Верни place для Cloud с именем "box".
func task2() string {
	return ""
}

// Задание 3
// Опиши type Fail struct с методом Error() string.
// Метод возвращает "сбой".
// Положи Fail{} в переменную типа error и верни текст Error().
func task3() string {
	return ""
}
