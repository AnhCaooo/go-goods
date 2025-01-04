package helpers

import (
	"bytes"
	"reflect"
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

func TestMapInterfaceToStruct(t *testing.T) {
	type SampleStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name    string
		input   interface{}
		want    *SampleStruct
		wantErr bool
	}{
		{
			name: "valid input",
			input: map[string]interface{}{
				"name":  "test",
				"value": 123,
			},
			want: &SampleStruct{
				Name:  "test",
				Value: 123,
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			input: map[string]interface{}{
				"name":  "test",
				"value": "invalid",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := MapInterfaceToStruct[SampleStruct](test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("MapInterfaceToStruct() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !test.wantErr && !reflect.DeepEqual(got, test.want) {
				t.Errorf("MapInterfaceToStruct() = %v, want %v", got, test.want)
			}
		})
	}
}
