// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(10, 20, 30), 60; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(5), 5; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2([]int{10, 20}, []int{30, 40}), 100; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2([]int{1}, []int{2}), 3; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := 4
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
