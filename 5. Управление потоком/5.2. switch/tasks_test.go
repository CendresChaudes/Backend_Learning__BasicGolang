// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("GET"), "чтение"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1("POST"), "запись"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1("PUT"), "неизвестно"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(200), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(404), "missing"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(201), "unknown"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(6), "выходной"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(3), "работа"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(0), "нет"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
