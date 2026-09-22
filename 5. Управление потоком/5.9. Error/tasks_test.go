// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(""), "пустое имя"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1("Анна"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(10, 0), "деление на ноль"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(10, 2), "5"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(200), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(100), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(599), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(20), "плохой код"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(600), "плохой код"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
