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
		{Name: "bank", Secret: "c"},
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
// Верни индекс первой записи с таким Name.
// Если записи нет, верни -1.
func task2(items []Entry, name string) int {
	return 0
}

// Задание 3
// Убери из items все записи с таким Name.
// Верни Name оставшихся записей через запятую без пробела.
// Если ничего не осталось, верни пустую строку.
func task3(items []Entry, name string) string {
	return ""
}
