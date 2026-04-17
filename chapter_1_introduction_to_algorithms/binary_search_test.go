package binarysearch_test

import (
	binarysearch "study/grokking-algorithms/chapter_1_introduction_to_algorithms"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	got := binarysearch.BinarySearch(arr, 3)
	want := 2
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	// first
	got = binarysearch.BinarySearch(arr, 1)
	want = 0
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	// last
	got = binarysearch.BinarySearch(arr, 7)
	want = 6
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	//not found
	got = binarysearch.BinarySearch(arr, 20)
	want = -1
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	// empty slice
	arr = []int{}
	got = binarysearch.BinarySearch(arr, 7)
	want = -1
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
