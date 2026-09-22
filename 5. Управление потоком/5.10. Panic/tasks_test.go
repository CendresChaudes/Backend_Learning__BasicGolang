// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "стоп"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(-3), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(4), 4; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(0), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := true
	if got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
