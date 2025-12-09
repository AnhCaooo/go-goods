package testcontainer

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name       string
		image      Image
		shouldFail bool
	}{
		{
			name: "Valid MongoDB image",
			image: Image{
				Name: "mongo:latest",
				Port: "27017",
			},
			shouldFail: false,
		},
		{
			name: "Unsupported image",
			image: Image{
				Name: "redis:latest",
				Port: "6379",
			},
			shouldFail: true,
		},
		{
			name: "Invalid name",
			image: Image{
				Name: "notexistimage123abc",
				Port: "1234",
			},
			shouldFail: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			container, err := Setup(ctx, tc.image)

			if tc.shouldFail {
				if err == nil {
					t.Fatalf("expected error but got success: %+v", container)
				}
				return
			}

			// Should succeed
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if container == nil {
				t.Fatal("expected non-nil container")
			}

			// verify the URI format for MongoDB
			if !strings.HasPrefix(container.URI, "mongodb://") {
				t.Errorf("invalid MongoDB URI: %s", container.URI)
			}

			// ensure container object is not empty
			if container.Ctr == nil {
				t.Errorf("container Ctr should not be nil")
			}

			// cleanup
			if container.Ctr != nil {
				err := container.Ctr.Terminate(ctx)
				if err != nil {
					t.Errorf("failed to terminate container: %v", err)
				}
			}
		})
	}
}
