package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1())

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3("NO_SUCH"))
}

// Задание 1
// Верни значение переменной окружения APP_PORT через os.Getenv.
func task1() string {
	return os.Getenv("APP_PORT")
}

// Задание 2
// Прочитай файл app.env.
// Верни значение APP_PORT из этого файла.
// Комментарии и пустые строки пропускай.
// Если файла нет или имени нет, верни пустую строку.
func task2() string {
	return ""
}

// Задание 3
// Прочитай app.env и верни значение ключа key.
// Если ключа нет, верни пустую строку.
func task3(key string) string {
	return ""
}
