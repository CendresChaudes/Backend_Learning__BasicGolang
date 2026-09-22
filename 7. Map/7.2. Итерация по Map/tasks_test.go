// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(map[string]int{"ok": 200, "missing": 404}), 604; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(map[string]int{}), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(map[string]int{"ok": 200, "missing": 404, "down": 500}), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	codes := map[string]int{"ok": 200, "missing": 404}
	if got, want := task3(codes, "ok"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task3(codes, "other"), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
