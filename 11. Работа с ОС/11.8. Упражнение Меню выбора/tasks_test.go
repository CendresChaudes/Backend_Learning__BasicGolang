// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(0), "выход"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(2), "добавить"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(4), "удалить"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(9), "неизвестно"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(3), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(9), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3([]int{2, 3, 0, 1}), "добавить найти"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3([]int{0, 1}), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3([]int{1, 4}), "показать удалить"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
