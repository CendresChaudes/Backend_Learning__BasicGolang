// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "weather"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "https://api.weather.local/v1"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := "weather https://api.weather.local/v1"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
