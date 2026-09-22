// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("vault"), "vault"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(7), "7"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(120), 3; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2("go"), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("vault"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task3(nil), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
