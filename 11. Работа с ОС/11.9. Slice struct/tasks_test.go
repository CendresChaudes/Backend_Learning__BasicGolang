// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := 2
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
	}
	if got, want := task2(items), "vpn"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(nil), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
		{Name: "mail", Secret: "c"},
	}
	if got, want := task3(items, "mail"), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task3(items, "bank"), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
