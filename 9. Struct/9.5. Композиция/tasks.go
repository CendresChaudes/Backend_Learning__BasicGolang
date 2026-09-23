package main

import "fmt"

type User struct {
	Name string
}

type Admin struct {
	User
	Level int
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
// Опиши метод Label() string у User: "user: " и имя.
// Создай Admin с именем "Ann" и уровнем 2.
// Верни adm.Name.
func (user User) Label() string {
	return "user: " + user.Name
}

func task1() string {
	return Admin{
		User:  User{Name: "Ann"},
		Level: 2,
	}.Name
}

// Задание 2
// У того же Admin верни Level.
func task2() int {
	return Admin{
		User:  User{Name: "Ann"},
		Level: 2,
	}.Level
}

// Задание 3
// У того же Admin верни результат Label().
func task3() string {
	return Admin{
		User:  User{Name: "Ann"},
		Level: 2,
	}.Label()
}
