// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		name string
		kind string
		want string
	}{
		{name: "ясно", kind: "clear", want: "ясно"},
		{name: "дождь", kind: "rain", want: "дождь"},
		{name: "снег", kind: "snow", want: "снег"},
		{name: "чужой код", kind: "fog", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task1(tt.kind); got != tt.want {
				t.Errorf("получилось %q, нужно %q", got, tt.want)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	tests := []struct {
		name string
		code int
		want string
	}{
		{name: "ok", code: 200, want: "ok"},
		{name: "missing", code: 404, want: "missing"},
		{name: "down", code: 500, want: "down"},
		{name: "unknown", code: 201, want: "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task2(tt.code); got != tt.want {
				t.Errorf("получилось %q, нужно %q", got, tt.want)
			}
		})
	}
}

func TestTask3(t *testing.T) {
	tests := []struct {
		name  string
		temp  int
		units string
		want  string
	}{
		{name: "цельсий", temp: 12, units: "C", want: "12 C"},
		{name: "фаренгейт", temp: 70, units: "F", want: "70 F"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task3(tt.temp, tt.units); got != tt.want {
				t.Errorf("получилось %q, нужно %q", got, tt.want)
			}
		})
	}
}
