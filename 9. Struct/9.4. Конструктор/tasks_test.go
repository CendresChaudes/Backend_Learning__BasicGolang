// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("Ann", 20), "Ann"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(), "пустое имя"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(), "отрицательный возраст"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
