// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := 200
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "ok"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(10, 3), 7; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(5, 1), 4; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
