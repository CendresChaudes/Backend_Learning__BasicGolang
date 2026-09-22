package main

import "fmt"

type Entry struct {
	Name   string
	Secret string
}

func main() {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
	}

	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2(items))

	fmt.Println("Задание 3")
	fmt.Println(task3(items, "mail"))
}

// Задание 1
// Собери срез из двух Entry: Name "mail" и Name "vpn".
// Верни длину среза.
func task1() int {
	return 0
}

// Задание 2
// Верни Name последней записи.
// Если срез пустой, верни пустую строку.
func task2(items []Entry) string {
	return ""
}

// Задание 3
// Верни, сколько записей имеют Name, равный name.
func task3(items []Entry, name string) int {
	return 0
}
