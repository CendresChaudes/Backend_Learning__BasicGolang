// Запуск: go test .

package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestTask1(t *testing.T) {
	path := filepath.Join(t.TempDir(), "account.json")
	got := task1(path)
	want := `{"login":"ann","secret":"qwe"}`
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != want {
		t.Errorf("в файле получилось %q, нужно %q", string(data), want)
	}
}

func TestTask2(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "account.json")
	if got, want := task2(path, ""), "пустой логин"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("пустой логин не должен создавать файл")
	}

	if got, want := task2(path, "ann"), ""; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
	data, err := os.ReadFile(path)
	wantFile := `{"login":"ann"}`
	if err != nil || string(data) != wantFile {
		t.Errorf("в файле получилось %q, нужно %q", string(data), wantFile)
	}
}

func TestTask3(t *testing.T) {
	path := filepath.Join(t.TempDir(), "account.json")
	got := task3(path)
	wantText := `{"login":"ann","port":80}`
	data, err := os.ReadFile(path)
	if err != nil || string(data) != wantText || got != len(wantText) {
		t.Errorf("получилось %d, нужно записать %q и вернуть %d", got, wantText, len(wantText))
	}
}
