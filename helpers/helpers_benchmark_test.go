package helpers

import "testing"

func BenchmarkTrimSpaceForByte(b *testing.B) {
	for b.Loop() {
		TrimSpaceForByte([]byte("   Hello, World!   "))
	}
}

func BenchmarkMapInterfaceToStruct(b *testing.B) {
	type SampleStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	input := map[string]interface{}{
		"name":  "test",
		"value": 123,
	}

	for b.Loop() {
		MapInterfaceToStruct[SampleStruct](input)
	}
}
