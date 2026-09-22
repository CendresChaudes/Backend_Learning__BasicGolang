// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1([]int{0, 1, 2, 3, 4}, 2), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1([]int{0, 1, 2}, 9), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	names := []string{"выход", "показать", "добавить", "найти", "удалить"}
	if got, want := task2(names, "добавить"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(names, "нет"), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	names := []string{"выход", "показать", "добавить", "найти", "удалить"}
	if got, want := task3(names, "найти"), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(names, "нет"), "нет"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
