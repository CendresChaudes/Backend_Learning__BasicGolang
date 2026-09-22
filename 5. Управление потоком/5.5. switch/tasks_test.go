// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("GET"), "чтение"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1("POST"), "создание"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1("PUT"), "неизвестно"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(200), "успех"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(201), "успех"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(404), "нет"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(500), "другое"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(1), "будни"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(5), "будни"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(6), "выходные"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(7), "выходные"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(0), "нет"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
