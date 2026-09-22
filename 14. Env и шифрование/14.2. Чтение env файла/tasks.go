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
// Опиши func envValue(text, key string) string.
// Она ищет строку key=значение и возвращает значение.
// Пустые строки и строки, которые начинаются с #, пропускай.
// Прочитай файл app.env через os.ReadFile.
// Верни значение ключа VAULT_KEY.
// Если файл не прочитался, верни пустую строку.
func task1() string {
	return ""
}

// Задание 2
// Прочитай app.env и верни значение ключа APP_PORT.
// Для разбора текста используй envValue.
func task2() string {
	return ""
}

// Задание 3
// Прочитай app.env и верни значение ключа APP_MODE.
func task3() string {
	return ""
}
