package golangsorter

func partition(arr []int, min int, max int) int {
	pivot := arr[max]
	i := min - 1
	for j := min; j <= max-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[max] = arr[max], arr[i+1]
	return i + 1
}

func quickSort(arr []int, min int, max int) {
	if min < max {
		pi := partition(arr, min, max)

		quickSort(arr, min, pi-1)
		quickSort(arr, pi+1, max)
	}
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
