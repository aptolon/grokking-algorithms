package recursion

import "fmt"

func Сountdown(i int) {
	fmt.Println(i)
	if i > 1 {
		Сountdown(i - 1)
	}
}

func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * (Factorial(i - 1))
}

func SumArr(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + SumArr(arr[1:])
}

func MaxNum(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}
	subMax := MaxNum(arr[1:])
	if subMax > arr[0] {
		return subMax
	}
	return arr[0]
}
