package recursion_test

import (
	"testing"

	recursion "study/grokking-algorithms/chapter_3_recursion"
)

func TestFactorial(t *testing.T) {
	// 0
	got := recursion.Factorial(0)
	want := 1
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// обычный случай
	got = recursion.Factorial(5)
	want = 120
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// ещё случай
	got = recursion.Factorial(3)
	want = 6
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSumArr(t *testing.T) {
	// обычный случай
	arr := []int{1, 2, 3, 4}
	got := recursion.SumArr(arr)
	want := 10
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// один элемент
	arr = []int{42}
	got = recursion.SumArr(arr)
	want = 42
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// пустой массив
	arr = []int{}
	got = recursion.SumArr(arr)
	want = -1
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// с отрицательными числами
	arr = []int{-1, 2, -3, 4}
	got = recursion.SumArr(arr)
	want = 2
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMaxNum(t *testing.T) {
	// обычный случай
	arr := []int{1, 5, 3, 9, 2}
	got := recursion.MaxNum(arr)
	want := 9
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// один элемент
	arr = []int{42}
	got = recursion.MaxNum(arr)
	want = 42
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	// пустой массив
	arr = []int{}
	got = recursion.MaxNum(arr)
	want = -1
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// отрицательные числа
	arr = []int{-10, -3, -50, -1}
	got = recursion.MaxNum(arr)
	want = -1
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// уже отсортирован
	arr = []int{1, 2, 3, 4, 5}
	got = recursion.MaxNum(arr)
	want = 5
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	// в обратном порядке
	arr = []int{5, 4, 3, 2, 1}
	got = recursion.MaxNum(arr)
	want = 5
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
