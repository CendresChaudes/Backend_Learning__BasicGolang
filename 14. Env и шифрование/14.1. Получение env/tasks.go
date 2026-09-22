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
// Верни os.Getenv("APP_PORT").
func task1() string {
	return ""
}

// Задание 2
// Прочитай APP_MODE через os.Getenv.
// Если строка пустая, верни "dev".
// Если в переменной есть текст, верни его.
func task2() string {
	return ""
}

// Задание 3
// Верни true, если переменная VAULT_KEY задана.
// os.LookupEnv возвращает значение и ok.
// ok равен true, когда переменная есть.
func task3() bool {
	return false
}
