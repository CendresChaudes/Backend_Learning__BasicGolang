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
// Опиши Account: Login с тегом json "login", Secret с тегом json "secret".
// Собери JSON для Login "kate" и Secret "xyz". Верни текст.
func task1() string {
	return ""
}

// Задание 2
// Опиши отдельный тип Session: Login с тегом json "login", Code с тегом json "-".
// Заполни Login "bob" и Code "123". Верни JSON.
// Поле Code в текст попасть не должно.
func task2() string {
	return ""
}

// Задание 3
// Опиши отдельный тип Profile: Login с тегом json "login",
// Nickname с тегом json "nickname,omitempty".
// Заполни Login "ivan" и оставь Nickname пустым. Верни JSON.
func task3() string {
	return ""
}
