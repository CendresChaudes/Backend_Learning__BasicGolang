// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1("")
	want := "C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task1("F")
	want = "F"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2("https://api.weather.local", "/v1")
	want := "https://api.weather.local/v1"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3(12, "C")
	want := "12 C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task3(0, "F")
	want = "0 F"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
