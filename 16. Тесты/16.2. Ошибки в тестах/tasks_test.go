// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(""), "пустой город"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1("Moscow"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(404), "плохой статус"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(200), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(""), "пустое тело"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3("ясно"), "ясно"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
