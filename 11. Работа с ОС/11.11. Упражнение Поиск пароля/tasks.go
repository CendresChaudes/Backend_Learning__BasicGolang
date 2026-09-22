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
	fmt.Println(task1(items, "vpn"))

	fmt.Println("Задание 2")
	fmt.Println(task2(items, "vpn"))

	fmt.Println("Задание 3")
	fmt.Println(task3(items, "vpn"))
}

// Задание 1
// Верни Secret первой записи с Name, равным name.
// Если такой записи нет, верни пустую строку.
func task1(items []Entry, name string) string {
	return ""
}

// Задание 2
// Верни true, если в items есть запись с таким Name.
func task2(items []Entry, name string) bool {
	return false
}

// Задание 3
// Верни индекс первой записи с таким Name.
// Если записи нет, верни -1.
func task3(items []Entry, name string) int {
	return 0
}
