// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(200), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task1(404), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(8080), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(80), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(1024), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("go", "Go"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task3("go", "go"), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
