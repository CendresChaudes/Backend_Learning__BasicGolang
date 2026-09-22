// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(200), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(404), "ошибка"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(20, true), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(18, false), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(15, true), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(95), "отлично"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(75), "хорошо"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(40), "ещё раз"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
