// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	if got, want := task1(200), "ok"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(404), "нет"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task1(500), "другое"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	if got, want := task2(95), "отлично"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(90), "отлично"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(70), "зачёт"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task2(40), "незачёт"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(3), "плюс"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(-2), "минус"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if got, want := task3(0), "ноль"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
