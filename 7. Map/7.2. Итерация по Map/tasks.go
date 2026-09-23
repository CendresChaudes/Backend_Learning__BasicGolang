package main

import "fmt"

func main() {
	codes := map[string]int{"ok": 200, "missing": 404}
	fmt.Println("Задание 1")
	fmt.Println(task1(codes))

	fmt.Println("Задание 2")
	fmt.Println(task2(codes))

	fmt.Println("Задание 3")
	fmt.Println(task3(codes, "ok"))
}

// Задание 1
// Верни сумму всех значений map.
func task1(codes map[string]int) int {
	var sum int
	for _, v := range codes {
		sum += v
	}

	return sum
}

// Задание 2
// Верни, сколько значений строго больше 300.
func task2(codes map[string]int) int {
	var count int
	for _, v := range codes {
		if v > 300 {
			count++
		}
	}

	return count
}

// Задание 3
// Верни true, если ключ key есть. Ищи его циклом range.
func task3(codes map[string]int, key string) bool {
	for k, _ := range codes {
		if k == key {
			return true
		}
	}

	return false
}
