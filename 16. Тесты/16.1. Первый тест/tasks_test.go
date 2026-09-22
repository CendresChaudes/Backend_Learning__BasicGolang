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
	if got, want := task2(404), "ошибка"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(4), 8; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}
