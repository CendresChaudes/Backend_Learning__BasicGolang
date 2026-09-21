// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "main"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2("backend"), "BACKEND"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2("go"), "GO"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("=", 20), "===================="; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3("-", 3), "---"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
