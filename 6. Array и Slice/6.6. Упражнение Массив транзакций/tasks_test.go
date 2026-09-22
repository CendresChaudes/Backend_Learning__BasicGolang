// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1([]int{100, -40, 25}), 85; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1([]int{10, -3}), 7; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2([]int{100, -40, -5, 25}), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2([]int{-1}), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2([]int{10, 20}), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3([]int{100, -40, 250, 25}), 250; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3([]int{3}), 3; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3([]int{-5, -1}), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
