// Запуск: go test .

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeVault(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "vault.json")
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatalf("не удалось подготовить файл: %v", err)
	}
	return path
}

func TestTask1(t *testing.T) {
	path := writeVault(t, `[{"name":"mail","secret":"a"},{"name":"vpn","secret":"b"}]`)
	got := task1(path)
	want := 2
	if got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	path := writeVault(t, `[{"name":"mail","secret":"a"},{"name":"vpn","secret":"b"}]`)
	if got, want := task2(path), "b"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no.json")
	if got, want := task3(missing), "нет файла"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	broken := writeVault(t, `{`)
	if got, want := task3(broken), "битый json"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	path := writeVault(t, `[{"name":"mail","secret":"a"}]`)
	if got, want := task3(path), "mail"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}
