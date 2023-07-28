package sorts

func ShellSort(arr []int) {
	length := len(arr)
	for gap := int(length / 2); gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			for j := i; j >= gap; j -= gap {
				if arr[j-gap] > arr[j] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
				}
			}
		}
	}
}
