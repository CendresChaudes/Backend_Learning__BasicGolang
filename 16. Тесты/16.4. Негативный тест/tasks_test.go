// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	got := task1("Moscow")
	want := "city=Moscow"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task1("")
	want = ""
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(200), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(404), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if got, want := task2(500), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
}

func TestTask3(t *testing.T) {
	got := task3("Moscow", "ясно")
	want := "Moscow: ясно"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task3("", "ясно")
	want = "нет данных"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	got = task3("Moscow", "")
	want = "нет данных"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
