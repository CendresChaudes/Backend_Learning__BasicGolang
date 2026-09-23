// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(1, 2, 3), 6; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(), 10; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3([]int{4, 5}), 9; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(nil), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
