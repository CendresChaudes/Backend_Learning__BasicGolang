// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1("Moscow", 12)
	want := "Moscow: 12 C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task1("Kazan", -5)
	want = "Kazan: -5 C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2("Moscow", "C")
	want := "city=Moscow&units=C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task2("New York", "F")
	want = "city=New+York&units=F"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3("Moscow")
	want := "https://api.weather.local/v1?city=Moscow&units=C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task3("SPb")
	want = "https://api.weather.local/v1?city=SPb&units=C"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
