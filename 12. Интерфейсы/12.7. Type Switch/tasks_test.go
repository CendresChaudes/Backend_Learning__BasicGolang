// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("vault"), "string"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(7), "int"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(true), "other"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2("vault"), 5; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(7), 7; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(true), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("go"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task3(7), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
