package helpers

import "testing"

func TestRemoveDuplicate_Ints(t *testing.T) {
	input := []int{1, 2, 2, 3, 1, 4}
	expected := []int{1, 2, 3, 4}
	result := RemoveDuplicate(input)
	if len(result) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("expected %v at index %d, got %v", v, i, result[i])
		}
	}
}

func TestRemoveDuplicate_Strings(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b"}
	expected := []string{"a", "b", "c"}
	result := RemoveDuplicate(input)
	if len(result) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("expected %v at index %d, got %v", v, i, result[i])
		}
	}
}

func TestRemoveDuplicate_Empty(t *testing.T) {
	input := []int{}
	result := RemoveDuplicate(input)
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestRemoveDuplicate_NoDuplicates(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}
	result := RemoveDuplicate(input)
	if len(result) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("expected %v at index %d, got %v", v, i, result[i])
		}
	}
}
