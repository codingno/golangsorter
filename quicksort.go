package golangsorter

func Partition(arr []int, min int, max int) int {
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

func QuickSort(arr []int, min int, max int) {
	if min < max {
		pi := Partition(arr, min, max)

		QuickSort(arr, min, pi-1)
		QuickSort(arr, pi+1, max)
	}
}
