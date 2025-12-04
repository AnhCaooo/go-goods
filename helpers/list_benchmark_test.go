package helpers

import "testing"

func BenchmarkRemoveDuplicate(b *testing.B) {
	for b.Loop() {
		input := []int{1, 2, 2, 3, 1, 4}
		RemoveDuplicate(input)
	}
}
