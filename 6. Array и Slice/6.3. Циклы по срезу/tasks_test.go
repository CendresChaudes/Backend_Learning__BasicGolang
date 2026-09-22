// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1([]int{200, 404, 500}), 1104; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(nil), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2([]int{100, 404, 200, 500}), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2([]int{1, 2}), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3([]int{200, 201, 404}, 404), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3([]int{200, 201}, 404), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
