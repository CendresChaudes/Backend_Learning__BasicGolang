// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "I can eat glass and it doesn't hurt me."
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "Don't communicate by sharing memory, share memory by communicating."
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := "If a program is too slow, it must have a loop."
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
