// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	items := []Account{
		{Login: "anna", Secret: "a1"},
		{Login: "boris", Secret: "b2"},
	}
	if got, want := task1(items, "boris"), "b2"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(items, "bank"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	items := []Account{
		{Login: "anna", Secret: "a1"},
		{Login: "boris", Secret: "b2"},
	}
	if got, want := task2(items, "anna"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(items, "bank"), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	items := []Account{
		{Login: "anna", Secret: "a1"},
		{Login: "boris", Secret: "b2"},
	}
	if got, want := task3(items, "boris"), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(items, "bank"), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
