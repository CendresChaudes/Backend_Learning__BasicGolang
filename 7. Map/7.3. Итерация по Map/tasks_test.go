// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(map[string]int{"tea": 10, "coffee": 20}), 30; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task1(map[string]int{"pen": 7}), 7; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(map[string]int{"a": 5, "b": 11, "c": 40}), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if got, want := task2(map[string]int{"a": 1}), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(map[string]int{"ok": 200, "missing": 404}), "missing"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(map[string]int{"ok": 200}), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
