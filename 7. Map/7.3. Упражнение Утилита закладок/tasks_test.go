// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	marks := map[string]string{"docs": "https://go.dev"}
	if got, want := task1(marks, "docs"), "https://go.dev"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(marks, "blog"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	marks := map[string]string{"docs": "https://go.dev"}
	if got, want := task2(marks, "docs"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(marks, "blog"), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(map[string]string{"docs": "https://go.dev"}, "blog", "https://go.dev/blog"), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	marks := map[string]string{"docs": "https://go.dev"}
	if got, want := task3(marks, "docs", "https://go.dev/doc"), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
