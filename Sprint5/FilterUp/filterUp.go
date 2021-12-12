package filterUp

func SiftUp(heap []int, idx int) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	for heap[idx] > heap[idx/2] {
		if idx == 1 {
			break
		}
		heap[idx], heap[idx/2] = heap[idx/2], heap[idx]
		idx /= 2

	}
	return idx
}

func SiftUpRecurcial(heap []int, idx int) int {
	if idx == 1 {
		return idx
	}
	parrent := idx / 2
	if heap[idx] > heap[parrent] {
		heap[idx], heap[parrent] = heap[parrent], heap[idx]
		return SiftUpRecurcial(heap, parrent)
	} else {
		return idx
	}
}
