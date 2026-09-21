// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "api"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "worker"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(8080), 8080; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(80), 80; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
