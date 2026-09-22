// Запуск: go test .

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTask1(t *testing.T) {
	path := filepath.Join(t.TempDir(), "note.txt")
	if err := os.WriteFile(path, []byte("секрет"), 0644); err != nil {
		t.Fatalf("не удалось подготовить файл: %v", err)
	}
	got := task1(path)
	want := "секрет"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask2(t *testing.T) {
	path := filepath.Join(t.TempDir(), "note.txt")
	if err := os.WriteFile(path, []byte("go"), 0644); err != nil {
		t.Fatalf("не удалось подготовить файл: %v", err)
	}
	got := task2(path)
	want := 2
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask3(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing.txt")
	if got, want := task3(missing), "нет файла"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	path := filepath.Join(t.TempDir(), "note.txt")
	if err := os.WriteFile(path, []byte("секрет"), 0644); err != nil {
		t.Fatalf("не удалось подготовить файл: %v", err)
	}
	if got, want := task3(path), "секрет"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
