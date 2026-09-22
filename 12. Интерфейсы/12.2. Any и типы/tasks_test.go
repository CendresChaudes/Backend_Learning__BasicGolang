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
	if got, want := task2("vault"), "vault"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(7), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(7), 7; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3("vault"), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
