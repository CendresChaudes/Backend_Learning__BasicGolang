// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "mail", Secret: "b"},
		{Name: "vpn", Secret: "c"},
	}
	if got, want := task1(items, "mail"), "a"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(items, "bank"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
	}
	if got, want := task2(items, "vpn"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(items, "bank"), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
	}
	if got, want := task3(items, "vpn"), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(items, "bank"), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
