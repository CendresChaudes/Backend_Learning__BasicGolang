// Запуск: go test .

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTask1(t *testing.T) {
	path := filepath.Join(t.TempDir(), "note.txt")
	if got, want := task1(path, "привет"), "привет"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("файл не записан: %v", err)
	}
	if string(data) != "привет" {
		t.Errorf("в файле %q, нужно %q", string(data), "привет")
	}
}

func TestTask2(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.txt")
	if got, want := task2(path), "нет файла"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if err := os.WriteFile(path, []byte("есть"), 0644); err != nil {
		t.Fatal(err)
	}
	if got, want := task2(path), "есть"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	if got, want := task3(), "открыл|закрыл"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
