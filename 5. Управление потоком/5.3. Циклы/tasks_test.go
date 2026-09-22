// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(3), 6; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(1), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(0), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(6), 12; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(1), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3([]int{1, 4, 15, 2}), 15; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3([]int{11, 20}), 11; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3([]int{1, 2, 3}), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
