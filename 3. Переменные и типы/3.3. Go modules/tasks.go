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
// Верни команду, которой создают модуль для программы по имени wallet.
func task1() string {
	return "go mod init wallet"
}

// Задание 2
// Верни две строки так, как они выглядят в go.mod для модуля wallet на Go 1.27.1.
// Между строками поставь \n.
func task2() string {
	return "module wallet\ngo 1.27.1"
}

// Задание 3
// Верни одно слово: что go.mod запоминает про чужие пакеты.
// Слово: зависимости.
func task3() string {
	return "зависимости"
}
