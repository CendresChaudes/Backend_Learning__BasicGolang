// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	t.Setenv("APP_PORT", "8080")
	got := task1()
	want := "8080"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	t.Setenv("APP_MODE", "")
	if got, want := task2(), "dev"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	t.Setenv("APP_MODE", "prod")
	if got, want := task2(), "prod"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	t.Setenv("VAULT_KEY", "1234567890123456")
	got := task3()
	want := true
	if got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
