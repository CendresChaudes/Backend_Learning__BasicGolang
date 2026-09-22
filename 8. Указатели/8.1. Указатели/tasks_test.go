// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(200), 200; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(200), 201; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	n := 10
	if got, want := task3(&n), 11; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if n != 11 {
		t.Errorf("переменная получилась %d, нужна 11", n)
	}
	if got, want := task3(nil), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
