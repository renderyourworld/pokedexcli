package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"   ", []string{}},
		{"hello", []string{"hello"}},
		{"  hello  ", []string{"hello"}},
		{"hello world", []string{"hello", "world"}},
		{"  hello   world  ", []string{"hello", "world"}},
		{"  Hello   World  ", []string{"hello", "world"}},
	}

	for _, test := range tests {
		result := cleanInput(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("cleanInput(%q) = %v; want %v", test.input, result, test.expected)
			break
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("cleanInput(%q) = %v; want %v", test.input, result, test.expected)
				break
			}
		}
	}
}
