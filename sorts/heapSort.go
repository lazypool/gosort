package sorts

func HeapSort(arr []int) {
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		sink(arr, i, length)
	}
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		sink(arr, 0, i)
	}
}

func sink(arr []int, index int, length int) {
	max := 0
	for left := index*2 + 1; left < length; left = index*2 + 1 {
		right := left + 1
		if right < length && arr[right] > arr[left] {
			max = right
		} else {
			max = left
		}
		if arr[index] < arr[max] {
			arr[index], arr[max] = arr[max], arr[index]
			index = max
		} else {
			break
		}
	}
}
