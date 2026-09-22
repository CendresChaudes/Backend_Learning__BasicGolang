// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("10 20 30"), 60; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1("4 8 15"), 27; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2("4 5 6 7"), 4; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2("9"), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("1 4 15 2"), 15; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3("11 20"), 11; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3("1 2 3"), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
