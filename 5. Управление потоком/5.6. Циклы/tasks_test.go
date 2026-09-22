// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(5), 15; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(1), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(4), 24; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(1), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(6), 3; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(1), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
