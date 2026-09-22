// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "zxy"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "aaa"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("abca", "abc"), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task3("abz", "abc"), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}
