package main

import "fmt"

type User struct {
	Name string
	Age  int
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
// Опиши func older(u User), которая ставит копии возраст 30.
// Создай User с возрастом 20, вызови older и верни возраст исходной переменной.
func task1() int {
	return 0
}

// Задание 2
// Опиши func grow(u *User), которая увеличивает возраст на 1.
// Создай User с возрастом 20, вызови grow от её адреса и верни новый возраст.
func task2() int {
	return 0
}

// Задание 3
// Опиши func rename(u *User, name string), которая пишет имя в исходный struct.
// Создай User, вызови rename(&user, "Mia") и верни имя.
func task3() string {
	return ""
}
