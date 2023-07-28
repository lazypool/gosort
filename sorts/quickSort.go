package sorts

func QuickSort(arr []int) {
	length := len(arr)
	handler(arr, 0, length)
}

func handler(arr []int, start, end int) {
	if end-start < 2 {
		return
	}
	middle := scanfR(arr, start, end)
	handler(arr, start, middle)
	handler(arr, middle+1, end)
}

func scanfR(arr []int, start, end int) int {
	for i, j := start, end-1; i < j; j-- {
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			return scanfL(arr, i, j+1)
		}
	}
	return start
}

func scanfL(arr []int, start, end int) int {
	for i, j := start, end-1; i < j; i++ {
		if arr[i] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
			return scanfR(arr, i, j+1)
		}
	}
	return end - 1
}
