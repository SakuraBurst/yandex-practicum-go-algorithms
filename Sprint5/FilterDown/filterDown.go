package filterDown

func SiftDown(heap []int, idx int) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	for {
		left := idx * 2
		rigth := idx*2 + 1
		if left >= len(heap) {
			break
		}
		if rigth < len(heap) {
			max := 0
			if heap[left] > heap[rigth] {
				max = left
			} else {
				max = rigth
			}
			if heap[max] > heap[idx] {
				heap[max], heap[idx] = heap[idx], heap[max]
				idx = max
			} else {
				break
			}
		} else {
			if heap[left] > heap[idx] {
				heap[left], heap[idx] = heap[idx], heap[left]
				idx = left
			} else {
				break
			}
		}
	}
	return idx
}

func SiftDown2(heap []int, idx int) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	for {
		left := idx * 2
		rigth := idx*2 + 1

		if left >= len(heap) {
			break
		}

		max := 0

		if rigth < len(heap) && heap[rigth] > heap[left] {
			max = rigth
		} else {
			max = left
		}

		if heap[max] > heap[idx] {
			heap[max], heap[idx] = heap[idx], heap[max]
			idx = max
		} else {
			break
		}
	}
	return idx
}

func SiftDown2Recurcial(heap []int, idx int) int {
	left := idx * 2
	right := idx*2 + 1

	if left >= len(heap) {
		return idx
	}

	max := 0

	if right < len(heap) && heap[right] > heap[left] {
		max = right
	} else {
		max = left
	}

	if heap[max] > heap[idx] {
		heap[max], heap[idx] = heap[idx], heap[max]
		return SiftDown2Recurcial(heap, max)
	} else {
		return idx
	}
}
