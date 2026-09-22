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
// Сделай Reader из строки "ясно" через strings.NewReader.
// Прочитай его целиком через io.ReadAll.
// Если ReadAll вернул ошибку, верни пустую строку.
// Иначе верни прочитанный текст.
func task1() string {
	return ""
}

// Задание 2
// Сделай Reader из строки "12 C".
// Прочитай его через io.ReadAll.
// Если ReadAll вернул ошибку, верни 0.
// Иначе верни число байтов в прочитанном срезе.
func task2() int {
	return 0
}

// Задание 3
// Сделай Reader из строки "ok".
// Прочитай его через io.ReadAll.
// Верни true, если текст равен "ok".
// Если ReadAll вернул ошибку, верни false.
func task3() bool {
	return false
}
