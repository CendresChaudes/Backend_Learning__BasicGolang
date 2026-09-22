// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "Moscow"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := "Moscow:F"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
