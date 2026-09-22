// Запуск: go test .

package main

import "testing"

func TestTask1(t *testing.T) {
	tests := []struct {
		name string
		kind string
		want string
	}{
		{name: "clear", kind: "clear", want: "ясно"},
		{name: "rain", kind: "rain", want: "дождь"},
		{name: "snow", kind: "snow", want: "снег"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := task1(tt.kind)
			if got != tt.want {
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
		{name: "not found", code: 404, want: "missing"},
		{name: "error", code: 500, want: "down"},
		{name: "other", code: 201, want: "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := task2(tt.code)
			if got != tt.want {
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
		{name: "celsius", temp: 12, units: "C", want: "12 C"},
		{name: "frost", temp: -3, units: "C", want: "-3 C"},
		{name: "fahrenheit", temp: 70, units: "F", want: "70 F"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := task3(tt.temp, tt.units)
			if got != tt.want {
				t.Errorf("получилось %q, нужно %q", got, tt.want)
			}
		})
	}
}
