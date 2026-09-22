package main

import "fmt"

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Опиши func apply(n int, op func(int) int) int. Она возвращает op(n).
// Опиши func double(n int) int. Она возвращает n * 2.
// В task1 верни apply(5, double).
func task1() int {
	return 0
}

// Задание 2
// Опиши func greet(name string, format func(string) string) string.
// Она возвращает format(name).
// Опиши func userLabel(name string) string. Она возвращает "user:" + name.
// В task2 верни greet("anna", userLabel).
func task2() string {
	return ""
}

// Задание 3
// Опиши func both(login string, a, b func(string) bool) bool.
// Она возвращает true, только когда a(login) и b(login) оба равны true.
// Опиши func isAdmin(login string) bool. Она возвращает login == "admin".
// Опиши func longEnough(login string) bool. Она возвращает len(login) > 3.
// В task3 верни both("admin", isAdmin, longEnough).
func task3() bool {
	return false
}
