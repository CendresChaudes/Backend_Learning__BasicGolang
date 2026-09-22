// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(100, []int{50, -30, -20}), 100; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(0, []int{10, -3}), 7; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2([]int{-10, 5, -3}), 13; got != want {
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
	if got, want := task3([]int{10, -4}), 6; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3([]int{-2, 5}), -2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(nil), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
