// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(true, true), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task1(true, false), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task1(false, true), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(200), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(201), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(404), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(false), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task3(true), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
