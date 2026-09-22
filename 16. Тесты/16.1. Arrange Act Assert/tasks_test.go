// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	city := "Moscow"
	text := "ясно"
	want := "Moscow: ясно"
	got := task1(city, text)
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	code := 200
	want := "ok"
	got := task2(code)
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	code = 404
	want = "fail"
	got = task2(code)
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := "ясно"
	want := "ясно"
	same := task3(got, want)
	if same != true {
		t.Errorf("получилось %t, нужно %t", same, true)
	}

	got = "дождь"
	want = "ясно"
	same = task3(got, want)
	if same != false {
		t.Errorf("получилось %t, нужно %t", same, false)
	}
}
