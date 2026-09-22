// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1("ясно", "дождь")
	want := "получилось \"ясно\", нужно \"дождь\""
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2("TestTask1", true)
	want := "PASS TestTask1"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task2("TestTask1", false)
	want = "FAIL TestTask1"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	rawURL := "https://api.weather.local/v1?city=Moscow"
	got := task3(200, rawURL)
	want := "ok"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task3(404, rawURL)
	want = "GET https://api.weather.local/v1?city=Moscow -> 404"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
