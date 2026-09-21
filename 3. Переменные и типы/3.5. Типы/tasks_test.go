// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "7 2.5 Go true"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "0 0.0 \"\" false"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(4, 2.5), 10.0; got != want {
		t.Errorf("получилось %v, нужно %v", got, want)
	}
	if got, want := task3(2, 1.5), 3.0; got != want {
		t.Errorf("получилось %v, нужно %v", got, want)
	}
}
