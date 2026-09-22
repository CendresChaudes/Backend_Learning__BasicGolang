// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	codes := map[string]int{"ok": 200, "missing": 404}
	if got, want := task1(codes, "ok"), 200; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	codes := map[string]int{"ok": 200, "zero": 0}
	if got, want := task2(codes, "other"), -1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(codes, "zero"), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(map[string]int{"ok": 200, "missing": 404}, "missing"), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
