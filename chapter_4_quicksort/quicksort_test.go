package quicksort_test

import (
	"reflect"
	"testing"

	quicksort "study/grokking-algorithms/chapter_4_quicksort"
)

func TestQuicksort(t *testing.T) {
	// обычный случай
	arr := []int{3, 1, 4, 1, 5, 9}
	got := quicksort.Quicksort(arr)
	want := []int{1, 1, 3, 4, 5, 9}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// пустой массив
	arr = []int{}
	got = quicksort.Quicksort(arr)
	want = []int{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// один элемент
	arr = []int{42}
	got = quicksort.Quicksort(arr)
	want = []int{42}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// уже отсортирован
	arr = []int{1, 2, 3, 4}
	got = quicksort.Quicksort(arr)
	want = []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// с отрицательными числами
	arr = []int{-3, 0, -1, 5, 2}
	got = quicksort.Quicksort(arr)
	want = []int{-3, -1, 0, 2, 5}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMergeSort(t *testing.T) {
	// обычный случай
	arr := []int{5, 2, 8, 1, 3}
	got := quicksort.MergeSort(arr)
	want := []int{1, 2, 3, 5, 8}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// пустой массив
	arr = []int{}
	got = quicksort.MergeSort(arr)
	want = []int{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// один элемент
	arr = []int{7}
	got = quicksort.MergeSort(arr)
	want = []int{7}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// уже отсортирован
	arr = []int{1, 2, 3, 4}
	got = quicksort.MergeSort(arr)
	want = []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// с отрицательными
	arr = []int{-2, -5, 0, 3}
	got = quicksort.MergeSort(arr)
	want = []int{-5, -2, 0, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestQuickSortInPlace(t *testing.T) {
	// обычный случай
	arr := []int{4, 2, 7, 1, 3}
	quicksort.QuickSortInPlace(arr)
	want := []int{1, 2, 3, 4, 7}

	if !reflect.DeepEqual(arr, want) {
		t.Fatalf("got %v, want %v", arr, want)
	}

	// пустой массив
	arr = []int{}
	quicksort.QuickSortInPlace(arr)
	want = []int{}

	if !reflect.DeepEqual(arr, want) {
		t.Fatalf("got %v, want %v", arr, want)
	}

	// один элемент
	arr = []int{10}
	quicksort.QuickSortInPlace(arr)
	want = []int{10}

	if !reflect.DeepEqual(arr, want) {
		t.Fatalf("got %v, want %v", arr, want)
	}

	// с повторяющимися элементами
	arr = []int{3, 3, 3, 1, 2}
	quicksort.QuickSortInPlace(arr)
	want = []int{1, 2, 3, 3, 3}

	if !reflect.DeepEqual(arr, want) {
		t.Fatalf("got %v, want %v", arr, want)
	}
}
