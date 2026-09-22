// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := true
	if got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := 7
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := "disk"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
