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
// Опиши type Wallet struct с полем Balance int.
// На уровне пакета напиши функцию: она принимает Wallet, прибавляет 15 к Balance копии
// и возвращает Balance этой копии.
// Начни с баланса 100. Верни то, что вернула функция.
func task1() int {
	return 0
}

// Задание 2
// Опиши type Counter struct с полем N int.
// На уровне пакета напиши функцию: она принимает Counter и ставит у копии N = 100.
// Создай счётчик с N = 7 и передай его в эту функцию.
// Верни N исходного счётчика.
func task2() int {
	return 0
}

// Задание 3
// Опиши type City struct с полем Name string.
// Создай два значения с Name "Riga".
// Верни true, если они равны через ==.
func task3() bool {
	return false
}
