package helpers

import (
	"bytes"
	"testing"
)

// TestTrimSpaceForByte tests the TrimSpaceForByte function.
func TestTrimSpaceForByte(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "No spaces",
			input:    []byte("HelloWorld"),
			expected: []byte("HelloWorld"),
		},
		{
			name:     "Leading spaces",
			input:    []byte("   HelloWorld"),
			expected: []byte("HelloWorld"),
		},
		{
			name:     "Trailing spaces",
			input:    []byte("HelloWorld   "),
			expected: []byte("HelloWorld"),
		},
		{
			name:     "Leading and trailing spaces",
			input:    []byte("   HelloWorld   "),
			expected: []byte("HelloWorld"),
		},
		{
			name:     "Only spaces",
			input:    []byte("     "),
			expected: []byte(""),
		},
		{
			name:     "Empty input",
			input:    []byte(""),
			expected: []byte(""),
		},
		{
			name:     "Spaces between words",
			input:    []byte("   Hello   World   "),
			expected: []byte("Hello   World"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TrimSpaceForByte(test.input)
			if !bytes.Equal(result, test.expected) {
				t.Errorf("TrimSpaceForByte(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}
