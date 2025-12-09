package testcontainer

import (
	"context"
	"testing"
)

func BenchmarkSetupTestContainer(b *testing.B) {
	ctx := context.Background()

	for b.Loop() {
		Setup(ctx, MongoDB)
	}
}
