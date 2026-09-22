// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	t.Setenv("VAULT_KEY", "1234567890123456")
	got := task1()
	want := "3031323334353637383961626921247a00f5e1224ea100706d6d7699c2a9"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	t.Setenv("VAULT_KEY", "1234567890123456")
	if got, want := task2(), "b2"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	t.Setenv("VAULT_KEY", "1234567890123456")
	if got, want := task3(), "mail"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
