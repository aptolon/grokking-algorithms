package selectionsort_test

import (
	"reflect"
	"testing"

	selectionsort "study/grokking-algorithms/chapter_2_selection_sort"
)

func TestSelectionSort(t *testing.T) {
	// обычный случай
	arr := []int{64, 25, 12, 22, 11}
	got := selectionsort.SelectionSort(arr)
	want := []int{11, 12, 22, 25, 64}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// уже отсортирован
	arr = []int{1, 2, 3, 4, 5}
	got = selectionsort.SelectionSort(arr)
	want = []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// в обратном порядке
	arr = []int{5, 4, 3, 2, 1}
	got = selectionsort.SelectionSort(arr)
	want = []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// с дубликатами
	arr = []int{3, 1, 2, 3, 1}
	got = selectionsort.SelectionSort(arr)
	want = []int{1, 1, 2, 3, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// один элемент
	arr = []int{42}
	got = selectionsort.SelectionSort(arr)
	want = []int{42}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	// пустой массив
	arr = []int{}
	got = selectionsort.SelectionSort(arr)
	want = []int{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
