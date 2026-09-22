// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1("")
	want := "пустой город"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task1("Moscow")
	want = ""
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := task2(404)
	want := "плохой статус"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task2(200)
	want = "ok"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3("")
	want := "пустое тело"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task3("ясно")
	want = "ясно"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
