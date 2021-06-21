package model

//user model
type ArrayToSort struct {
	Things []int
}

func Sort(tosort []int) {
	quickSort(tosort, 0, len(tosort))
}

func quickSort(arr []int, begin, end int) {
	if begin < end {
		partitionIndex := partition(arr, begin, end)
		// Recursively sort elements of the 2 sub-arrays
		quickSort(arr, begin, partitionIndex-1)
		quickSort(arr, partitionIndex+1, end)
	}
}

func partition(arr []int, begin, end int) int {
	pivot := arr[end]
	i := (begin - 1)
	for j := begin; j < end; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]

	return i + 1
}
