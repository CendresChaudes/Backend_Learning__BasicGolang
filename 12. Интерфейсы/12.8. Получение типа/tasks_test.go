// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("vault"), "vault"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(5), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(8), 8; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2("x"), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("go"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task3(1), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
