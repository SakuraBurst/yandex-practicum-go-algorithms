package maxPathInTree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Solution(root *Node) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	bestPath, bestLocalPath := SolutionCheat(root)
	if bestLocalPath > bestPath {
		return bestLocalPath
	}
	return bestPath
}

func SolutionCheat(root *Node) (int, int) {
	if root.Left == nil && root.Right == nil {
		return root.Value, root.Value
	}
	if root.Left != nil && root.Right != nil {
		bl, lbpl := SolutionCheat(root.Left)
		br, lbpr := SolutionCheat(root.Right)
		previeousBestPath := 0
		if lbpr > lbpl {
			previeousBestPath = lbpr
		} else {
			previeousBestPath = lbpl
		}
		bestPath := 0
		if br > bl {
			bestPath = br + root.Value
		} else {
			bestPath = bl + root.Value
		}

		if root.Value > bestPath {
			bestPath = root.Value
		}

		localBestPath := root.Value + br + bl

		if bestPath > localBestPath {
			localBestPath = bestPath
		}

		if localBestPath > previeousBestPath {
			return bestPath, localBestPath
		} else {
			return bestPath, previeousBestPath
		}
	}
	if root.Left != nil {
		bl, lbpl := SolutionCheat(root.Left)
		currentPath := bl + root.Value

		if root.Value > currentPath {
			currentPath = root.Value
		}

		if currentPath > lbpl {
			return currentPath, currentPath
		}
		return currentPath, lbpl
	}

	br, lbpr := SolutionCheat(root.Right)
	currentPath := br + root.Value

	if root.Value > currentPath {
		currentPath = root.Value
	}

	if currentPath > lbpr {
		return currentPath, currentPath
	}
	return currentPath, lbpr

}
