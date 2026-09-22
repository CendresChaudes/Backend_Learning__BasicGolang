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
// Метод Title возвращает Name.
// Опиши функцию title(box FileBox) string. Она возвращает box.Title()
// и не создаёт FileBox внутри.
// В task1 создай FileBox с именем "vault.json" и верни title от него.
func task1() string {
	return ""
}

// Задание 2
// Опиши функцию mark(box FileBox) string.
// Она возвращает "file " и сразу за ней box.Name.
// FileBox внутри mark не создавай.
// В task2 создай FileBox с именем "db.json" и верни mark от него.
func task2() string {
	return ""
}

// Задание 3
// Опиши функцию ready(box FileBox) bool.
// Она возвращает true, если box.Name не пустой.
// В task3 создай FileBox с именем "db.json" и верни ready от него.
func task3() bool {
	return false
}
