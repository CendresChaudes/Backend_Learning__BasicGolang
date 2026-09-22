// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1("ясно"), "ясно"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(), "ясно"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3("Moscow", "C"), "https://api.weather.local/v1?city=Moscow&units=C"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
