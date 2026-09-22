// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(""), "пустой порт"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1("8080"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(0), "деление на ноль"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(2), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(-1), "отрицательное"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(3), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
