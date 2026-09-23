package main

import "fmt"

func main() {
	codes := map[string]int{"ok": 200, "missing": 404}
	fmt.Println("Задание 1")
	fmt.Println(task1(codes, "ok"))

	fmt.Println("Задание 2")
	fmt.Println(task2(codes, "other"))

	fmt.Println("Задание 3")
	fmt.Println(task3(map[string]int{"ok": 200, "missing": 404}, "missing"))
}

// Задание 1
// Верни значение по ключу key.
func task1(codes map[string]int, key string) int {
	return codes[key]
}

// Задание 2
// Если ключ key есть, верни его значение.
// Если ключа нет, верни -1.
func task2(codes map[string]int, key string) int {
	value, ok := codes[key]
	if ok {
		return value
	}

	return -1
}

// Задание 3
// Удали ключ key и верни новую длину map.
func task3(codes map[string]int, key string) int {
	delete(codes, key)

	return len(codes)
}
