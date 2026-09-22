// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1([]int{1, 2, 3, 4}), 10; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1([]int{8}), 8; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2([]int{200, 404, 500}), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2([]int{404}), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2([]int{200, 200}), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3([]int{200, 404}), "200-404"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3([]int{500}), "500"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
