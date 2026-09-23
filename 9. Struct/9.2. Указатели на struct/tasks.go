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
func older(u User) User {
	u.Age = 30
	return u
}

func task1() int {
	user := User{Age: 20}
	older(user)

	return user.Age
}

// Задание 2
// Опиши func grow(u *User), которая увеличивает возраст на 1.
// Создай User с возрастом 20, вызови grow от её адреса и верни новый возраст.
func grow(u *User) {
	u.Age = u.Age + 1
}

func task2() int {
	user := User{Age: 20}
	grow(&user)

	return user.Age
}

// Задание 3
// Опиши func rename(u *User, name string), которая пишет имя в исходный struct.
// Создай User, вызови rename(&user, "Mia") и верни имя.
func rename(u *User, name string) {
	u.Name = name
}

func task3() string {
	user := User{}
	rename(&user, "Mia")

	return user.Name
}
