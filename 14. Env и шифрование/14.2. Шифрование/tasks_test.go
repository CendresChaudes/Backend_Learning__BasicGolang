// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "3031323334353637383961627d720c0df0d0d984191548689c68bd324d0c89e005"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(), "vault"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(), "короткие данные"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
