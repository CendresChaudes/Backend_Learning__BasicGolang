// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(4), 6; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(6), 15; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(5), 9; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(1), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(5), 40; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(3), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
