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
	got := task2()
	want := "3031323334353637383961626672100d8bc56da5460662a58605ff7169b89b52"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := "3031323334353637383961626921247a00f5e1224ea100706d6d7699c2a9"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
