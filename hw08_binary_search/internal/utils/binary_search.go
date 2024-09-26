package utils

func SearchValueFromSlice(arr []int, value int) int {
	if len(arr) == 0 {
		return -1
	}
	start, end := 0, len(arr)-1
	result := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == value {
			return mid
		}
		if arr[mid] < value {
			start = mid + 1
		} else {

			end = mid - 1
		}
	}
	return result
}
