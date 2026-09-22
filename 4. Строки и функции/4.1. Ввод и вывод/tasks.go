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
// Собери текст так, как его напечатал бы fmt.Println("Готово").
// Возьми fmt.Sprintln и верни эту строку.
func task1() string {
	return fmt.Sprintln("Готово")
}

// Задание 2
// Собери текст так, как его напечатал бы fmt.Print("id", 7).
// Возьми fmt.Sprint и верни эту строку. Перевода строки в конце нет.
func task2() string {
	return fmt.Sprint("id ", 7)
}

// Задание 3
// Прочитай число из строки "15" через fmt.Sscan.
// Запиши его в переменную и верни эту переменную.
func task3() int {
	var someVar int
	fmt.Sscan("15", &someVar)
	return someVar
}
