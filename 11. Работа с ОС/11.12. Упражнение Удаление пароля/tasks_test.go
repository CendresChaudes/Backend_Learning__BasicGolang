// Запуск: go test .

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestTask1(t *testing.T) {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
		{Name: "bank", Secret: "c"},
	}
	if got, want := task1(items, "vpn"), 2; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}

	duplicates := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
		{Name: "vpn", Secret: "c"},
	}
	if got, want := task1(duplicates, "vpn"), 1; got != want {
		t.Errorf("получилось %d, нужно %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	items := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
		{Name: "bank", Secret: "c"},
	}
	if got, want := task2(items, "vpn"), "mail,bank"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	full := []Entry{
		{Name: "mail", Secret: "a"},
		{Name: "vpn", Secret: "b"},
		{Name: "bank", Secret: "c"},
	}
	if got, want := task2(full, "missing"), "mail,vpn,bank"; got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}
}

func TestTask3(t *testing.T) {
	path := filepath.Join(t.TempDir(), "vault.json")
	initial := `[{"name":"mail","secret":"a"},{"name":"vpn","secret":"b"},{"name":"bank","secret":"c"}]`
	if err := os.WriteFile(path, []byte(initial), 0644); err != nil {
		t.Fatalf("не удалось подготовить файл: %v", err)
	}

	got := task3(path, "vpn")
	want := "mail,bank"
	if got != want {
		t.Errorf("получилось %q, нужно %q", got, want)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("файл не прочитан: %v", err)
		return
	}
	var list []Entry
	if err := json.Unmarshal(data, &list); err != nil {
		t.Errorf("в файле битый json: %v", err)
		return
	}
	if len(list) != 2 || list[0].Name != "mail" || list[1].Name != "bank" {
		t.Errorf("в файле получилось %q, нужны mail и bank", string(data))
	}
}
