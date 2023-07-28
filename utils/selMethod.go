package utils

import (
	"fmt"
	"gosort/sorts"
	"os"
)

func SelMethod(name string) func(arr []int) {
	var method func(arr []int)
	switch name {
	case "bubble":
		fmt.Printf("Using Bubble Sort...")
		method = sorts.BubbleSort
	case "bucket":
		fmt.Printf("Using Bucket Sort...")
		method = sorts.BucketSort
	case "counting":
		fmt.Printf("Using Counting Sort...")
		method = sorts.CountingSort
	case "heap":
		fmt.Printf("Using Heap Sort...")
		method = sorts.HeapSort
	case "insertion":
		fmt.Printf("Using Insertion Sort...")
		method = sorts.InsertionSort
	case "merge":
		fmt.Printf("Using Merge Sort...")
		method = sorts.MergeSort
	case "quick":
		fmt.Printf("Using Quick Sort...")
		method = sorts.QuickSort
	case "radix":
		fmt.Printf("Using Radix Sort...")
		method = sorts.RadixSort
	case "selection":
		fmt.Printf("Using Selection Sort...")
		method = sorts.SelectionSort
	case "shell":
		fmt.Printf("Using Shell Sort...")
		method = sorts.ShellSort
	default:
		fmt.Printf("Error indicator. No method named '" + name +
			"'.\nAvalable names are" +
			":\nbubble, bucket, counting, heap, insertion, merge, quick, radix, selection, shell" +
			".\nPlease use -h to acquire more detailed usage.\n")
		os.Exit(1)
	}
	return method
}
