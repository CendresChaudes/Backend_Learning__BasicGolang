// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	t.Setenv("APP_PORT", "8080")
	if got, want := task1(), "8080"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(), "8080"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("VAULT_KEY"), "1234567890123456"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3("NO_SUCH"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
