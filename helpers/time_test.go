package helpers

import (
	"testing"
)

func TestParseHour(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		expectError bool
	}{
		{"00:00", 0, false},
		{"07:00", 7, false},
		{"12:34", 12, false},
		{"23:59", 23, false},
		{"24:00", 0, true},   // invalid hour
		{"invalid", 0, true}, // completely invalid
		{"7:00", 7, false},   // single-digit hour
		{"07:0", 0, true},    // invalid minute format
	}

	for _, tt := range tests {
		hour, err := ParseHour(tt.input)
		if (err != nil) != tt.expectError {
			t.Errorf("ParseHour(%q) error = %v, expectError = %v", tt.input, err, tt.expectError)
			continue
		}
		if hour != tt.expected {
			t.Errorf("ParseHour(%q) = %d; want %d", tt.input, hour, tt.expected)
		}
	}
}

func TestParseHourMinute(t *testing.T) {
	tests := []struct {
		input       string
		wantHour    int
		wantMinute  int
		expectError bool
	}{
		{"00:00", 0, 0, false},
		{"07:00", 7, 0, false},
		{"07:15", 7, 15, false},
		{"09:30", 9, 30, false},
		{"23:59", 23, 59, false},
		{"24:00", 0, 0, true},
		{"invalid", 0, 0, true},
	}

	for _, tt := range tests {
		h, m, err := ParseHourMinute(tt.input)
		if (err != nil) != tt.expectError {
			t.Errorf("ParseHourMinute(%q) error = %v, expectError = %v", tt.input, err, tt.expectError)
			continue
		}
		if h != tt.wantHour || m != tt.wantMinute {
			t.Errorf("ParseHourMinute(%q) = %d:%d; want %d:%d", tt.input, h, m, tt.wantHour, tt.wantMinute)
		}
	}
}
