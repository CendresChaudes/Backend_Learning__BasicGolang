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
// На уровне пакета опиши type User struct с полями Name string и Age int.
// NewUser(name string, age int) (User, error) проверяет аргументы.
// Пустое имя — ошибка с текстом "пустое имя".
// Возраст меньше нуля — ошибка с текстом "возраст".
// Иначе верни заполненный User и nil.
// Из task1 верни true, если NewUser("Ann", 20) отработал без ошибки.
func task1() bool {
	return false
}

// Задание 2
// Вызови ту же NewUser с именем "" и возрастом 20.
// Верни текст ошибки.
func task2() string {
	return ""
}

// Задание 3
// Вызови ту же NewUser с именем "Ann" и возрастом -1.
// Верни текст ошибки.
func task3() string {
	return ""
}
