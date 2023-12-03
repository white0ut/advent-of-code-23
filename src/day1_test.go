package main

import "testing"

func TestFindFirst(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "one", want: 1},
		{input: "xone", want: 1},
		{input: "5two", want: 5},
		{input: "nine4", want: 9},
	}

	for _, tt := range tests {
		got := findFirst(tt.input)
		if got != tt.want {
			t.Fatalf("expected %d for %q, got %d", tt.want, tt.input, got)
		}
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "onetwo", want: 2},
		{input: "xone4five", want: 5},
		{input: "5two", want: 2},
		{input: "nine4", want: 4},
	}

	for _, tt := range tests {
		got := findLast(tt.input)
		if got != tt.want {
			t.Fatalf("expected %d for %q, got %d", tt.want, tt.input, got)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "two1nine", want: 29},
		{input: "eightwothree", want: 83},
		{input: "abcone2threexyz", want: 13},
		{input: "xtwone3four", want: 24},
		{input: "4nineeightseven2", want: 42},
		{input: "zoneight234", want: 14},
		{input: "7pqrstsixteen", want: 76},
	}

	for _, tt := range tests {
		got := partTwo(tt.input)
		if got != tt.want {
			t.Fatalf("expected %d for %q, got %d", tt.want, tt.input, got)
		}
	}
}
