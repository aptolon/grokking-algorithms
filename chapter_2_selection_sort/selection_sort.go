package selectionsort

func findMin(arr []int) int {
	minIdx := 0
	minItem := arr[0]

	for i := range arr {
		if arr[i] < minItem {
			minItem = arr[i]
			minIdx = i
		}
	}

	return minIdx
}
func remove(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func SelectionSort(arr []int) []int {

	sortedArr := make([]int, 0, len(arr))
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)

	for range len(arr) {
		idx := findMin(copyArr)
		sortedArr = append(sortedArr, copyArr[idx])
		copyArr = remove(copyArr, idx)
	}
	return sortedArr

}
