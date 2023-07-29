package sorts

func CountingSort(arr []int) {
	length := countLen(arr)
	counts := make([]int, length)
	for _, num := range arr {
		counts[num] += 1
	}
	sorted := 0
	for i := 0; i < length; i++ {
		for counts[i] > 0 {
			arr[sorted] = i
			sorted += 1
			counts[i] -= 1
		}
	}
}

func countLen(arr []int) int {
	var max int
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max + 1
}
