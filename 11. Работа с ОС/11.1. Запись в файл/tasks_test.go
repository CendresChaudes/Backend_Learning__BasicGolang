// Запуск: go test .

package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestTask1(t *testing.T) {
	dir := t.TempDir()
	emptyPath := filepath.Join(dir, "empty.txt")
	if got, want := task1(emptyPath, ""), 0; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
	if _, err := os.Stat(emptyPath); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("пустой текст не должен создавать файл")
	}

	path := filepath.Join(dir, "note.txt")
	got := task1(path, "go")
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "go" || got != 2 {
		t.Errorf("получилось %d, нужно записать %q и вернуть 2", got, "go")
	}
}

func TestTask2(t *testing.T) {
	path := filepath.Join(t.TempDir(), "perm.txt")
	got := task2(path)
	data, err := os.ReadFile(path)
	info, statErr := os.Stat(path)
	if err != nil || statErr != nil || string(data) != "ok" || got != 0644 || info.Mode().Perm() != 0644 {
		t.Errorf("получилось %d, нужно записать %q и вернуть 0644", got, "ok")
	}
}

func TestTask3(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "note.txt")
	if got, want := task3(path, ""), "пустой текст"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("пустой текст не должен создавать файл")
	}

	if got, want := task3(path, "go"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "go" {
		t.Errorf("в файле получилось %q, нужно %q", string(data), "go")
	}
}
