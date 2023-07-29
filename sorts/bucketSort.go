package sorts

func BucketSort(arr []int) {
	var min, max int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	buckets := make([][]int, 10)
	for _, v := range arr {
		i := ((v - min) / (max - min) * 9)
		buckets[i] = append(buckets[i], v)
	}
	j := 0
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			BubbleSort(bucket)
			copy(arr[j:], bucket)
			j += len(bucket)
		}
	}
}
