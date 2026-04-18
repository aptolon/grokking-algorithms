package quicksort

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	var less, greater, equal []int

	for _, val := range arr {
		if val < pivot {
			less = append(less, val)
		} else if val == pivot {
			equal = append(equal, val)
		} else {
			greater = append(greater, val)
		}
	}
	return append(append(Quicksort(less), equal...), Quicksort(greater)...)
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)

}

func merge(right, left []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(arr, low, high)
	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)

}
func QuickSortInPlace(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
