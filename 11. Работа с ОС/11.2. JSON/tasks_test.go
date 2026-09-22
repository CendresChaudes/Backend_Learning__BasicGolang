// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(), `{"login":"ann","secret":"qwe"}`; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(), `{"login":"ann","secret":"","nickname":"ani"}`; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	text := `[{"name":"mail","secret":"a"},{"name":"vpn","secret":"b"}]`
	if got, want := task3(text), "b"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3("[]"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
