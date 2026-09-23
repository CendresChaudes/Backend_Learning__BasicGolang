package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Задание 1")
	fmt.Println(task1("Ann", 20))

	fmt.Println("Задание 2")
	fmt.Println(task2())

	fmt.Println("Задание 3")
	fmt.Println(task3())
}

// Задание 1
// Опиши func NewUser(name string, age int) (User, error) на уровне пакета.
// Пустое имя — ошибка "пустое имя".
// Отрицательный возраст — ошибка "отрицательный возраст".
// Иначе верни заполненный User и nil.
// task1 вызывает NewUser. Если ошибка есть, верни её текст.
// Если ошибки нет, верни имя.

func NewUser(name string, age int) (User, error) {
	if name == "" {
		return User{}, errors.New("пустое имя")
	}

	if age < 0 {
		return User{}, errors.New("отрицательный возраст")
	}

	return User{Name: name, Age: age}, nil
}

func task1(name string, age int) string {
	user, err := NewUser(name, age)
	if err != nil {
		return err.Error()
	}

	return user.Name
}

// Задание 2
// Вызови NewUser с пустым именем и возрастом 20.
// Верни текст ошибки.
func task2() string {
	_, err := NewUser("", 20)
	return err.Error()
}

// Задание 3
// Вызови NewUser с именем "Ann" и возрастом -1.
// Верни текст ошибки.
func task3() string {
	_, err := NewUser("Ann", -1)
	return err.Error()
}
