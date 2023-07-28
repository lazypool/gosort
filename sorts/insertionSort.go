package sorts

func InsertionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
