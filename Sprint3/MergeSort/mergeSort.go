package mergeSort

func merge_sort(arr []int, lf int, rg int) {
	if rg-lf == 1 {
		return
	}
	middle := (lf + rg) >> 1
	merge_sort(arr, lf, middle)
	merge_sort(arr, middle, rg)
	a := merge(arr, lf, middle, rg)
	copy(arr[lf:rg], a)

}

func merge(arr []int, left int, mid int, right int) (result []int) {
	midCopy := mid
	result = make([]int, 0, right-left)
	for left < mid || midCopy < right {
		if left == mid {
			result = append(result, arr[midCopy])
			midCopy++
			continue
		}
		if midCopy == right {
			result = append(result, arr[left])
			left++
			continue
		}
		if arr[left] <= arr[midCopy] {
			result = append(result, arr[left])
			left++
		} else {
			result = append(result, arr[midCopy])
			midCopy++
		}
	}
	return result
}
