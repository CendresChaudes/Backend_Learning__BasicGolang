// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := 42
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := 8
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := 5
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
