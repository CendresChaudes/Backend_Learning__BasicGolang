// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := 10
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	n := 10
	if got, want := task2(&n), 15; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if n != 15 {
		t.Errorf("в переменной %d, нужно %d", n, 15)
	}

	n = 1
	if got, want := task2(&n), 6; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if n != 6 {
		t.Errorf("в переменной %d, нужно %d", n, 6)
	}
}

func TestTask3(t *testing.T) {
	a, b := 2, 9
	if got, want := task3(&a, &b), 9; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if b != 2 {
		t.Errorf("во второй переменной %d, нужно %d", b, 2)
	}

	a, b = 8, 1
	if got, want := task3(&a, &b), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if b != 8 {
		t.Errorf("во второй переменной %d, нужно %d", b, 8)
	}
}
