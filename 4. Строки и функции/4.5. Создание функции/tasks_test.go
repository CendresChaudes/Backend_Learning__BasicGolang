// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(20, 22), 42; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(1, 2), 3; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2("Казань")
	want := "город Казань"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(4), 8; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(3), 6; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
