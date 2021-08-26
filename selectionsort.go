package golangsorter

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min_index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
			arr[min_index], arr[i] = arr[i], arr[min_index]
		}
	}
}
