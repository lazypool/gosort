package sorts

func MergeSort(arr []int) {
	length := len(arr)
	divide(arr, 0, length)
}

func divide(arr []int, start, end int) {
	if end-start < 2 {
		return
	}
	middle := int((start + end) / 2)
	divide(arr, start, middle)
	divide(arr, middle, end)
	merge(arr, start, middle, end)
}

func merge(arr []int, start, middle, end int) {
	i, j, k := start, middle, end
	for i < j && j < k {
		for i < j && arr[i] <= arr[j] {
			i++
		}
		for j < k && arr[j] <= arr[i] {
			j++
		}
		rotate(arr, i, middle, j)
		i = i + j - middle
		middle = j
	}
}

func rotate(arr []int, start, middle, end int) {
	reverse(arr, start, middle)
	reverse(arr, middle, end)
	reverse(arr, start, end)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end-1] = arr[end-1], arr[start]
		start++
		end--
	}
}
