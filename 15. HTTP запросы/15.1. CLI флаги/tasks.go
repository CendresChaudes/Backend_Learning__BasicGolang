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
// Создай FlagSet с именем "weather" и flag.ContinueOnError.
// Опиши флаг city через String. Значение по умолчанию — пустая строка.
// Разбери аргументы []string{"-city", "Moscow"}.
// Если Parse вернул ошибку, верни пустую строку.
// Иначе верни значение флага city.
func task1() string {
	return ""
}

// Задание 2
// Создай FlagSet так же.
// Опиши флаг units через String. Значение по умолчанию — "C".
// Разбери пустой срез []string{}.
// Если Parse вернул ошибку, верни пустую строку.
// Иначе верни значение флага units.
func task2() string {
	return ""
}

// Задание 3
// Создай FlagSet так же.
// Опиши флаги city (по умолчанию "") и units (по умолчанию "C").
// Разбери []string{"-city", "Moscow", "-units", "F"}.
// Если Parse вернул ошибку, верни пустую строку.
// Иначе верни город, двоеточие и единицы одной строкой: "Moscow:F".
func task3() string {
	return ""
}
