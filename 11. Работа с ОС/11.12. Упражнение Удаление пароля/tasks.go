package main

import "fmt"

type Entry struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
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
	fmt.Println(task3("vault.json", "vpn"))
}

// Задание 1
// Убери из items все записи с таким Name.
// Верни длину оставшегося среза. Порядок остальных не меняй.
func task1(items []Entry, name string) int {
	return 0
}

// Задание 2
// Убери из items все записи с таким Name.
// Верни Name оставшихся записей через запятую без пробела.
// Если ничего не осталось, верни пустую строку.
func task2(items []Entry, name string) string {
	return ""
}

// Задание 3
// Прочитай JSON-массив Entry из path.
// Убери все записи с таким Name и запиши новый JSON в тот же path. Права 0644.
// Верни Name оставшихся записей через запятую без пробела.
func task3(path, name string) string {
	return ""
}
