// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	nums := [3]int{1, 2, 3}
	if got, want := task1(&nums), 3; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if nums != [3]int{3, 2, 1} {
		t.Errorf("массив получился %v, нужен %v", nums, [3]int{3, 2, 1})
	}
}

func TestTask2(t *testing.T) {
	nums := [4]int{10, 20, 30, 40}
	if got, want := task2(&nums), 10; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if nums != [4]int{40, 30, 20, 10} {
		t.Errorf("массив получился %v, нужен %v", nums, [4]int{40, 30, 20, 10})
	}
}

func TestTask3(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 9}
	if got, want := task3(&nums), true; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if nums != [5]int{9, 4, 3, 2, 1} {
		t.Errorf("массив получился %v, нужен %v", nums, [5]int{9, 4, 3, 2, 1})
	}

	nums = [5]int{9, 2, 3, 4, 1}
	if got, want := task3(&nums), false; got != want {
		t.Errorf("получилось %t, нужно %t", got, want)
	}
	if nums != [5]int{1, 4, 3, 2, 9} {
		t.Errorf("массив получился %v, нужен %v", nums, [5]int{1, 4, 3, 2, 9})
	}
}
