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
// Создай User с именем "Ann" и возрастом 20. Верни имя.
func task1() string {
	return User{Name: "Ann", Age: 20}.Name
}

// Задание 2
// Создай User с именем "Ann" и возрастом 20. Верни возраст.
func task2() int {
	return User{Name: "Ann", Age: 20}.Age
}

// Задание 3
// Создай User только с именем "Ann". Возраст не заполняй.
// Верни длину имени плюс возраст.
// Незаполненный возраст равен 0, поэтому для "Ann" получится 3.
func task3() int {

	user := User{Name: "Ann"}
	name := user.Name
	age := user.Age

	return len(name + fmt.Sprint(age))
}
