// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1()
	want := "GET https://api.weather.local/v1?city=Moscow -> 200"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2()
	want := "GET https://api.weather.local/v1?city=Moscow -> 404"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3()
	want := "Moscow: ясно"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
