package journey

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func Journey(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	attractionsData := strings.Fields(scanner.Text())
	attractionsRating := make([]int, 0, n)
	for _, attractionRatingString := range attractionsData {
		attractionRating, _ := strconv.Atoi(attractionRatingString)
		attractionsRating = append(attractionsRating, attractionRating)
	}
	dpPrev := make([]int, n+1)
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		maxPath := 0
		for g := 1; g < i+1; g++ {
			dp[g] = dpPrev[g]
			if attractionsRating[g-1] < attractionsRating[i-1] {
				maxPath = max(maxPath, dp[g])
			}
			if attractionsRating[g-1] == attractionsRating[i-1] {
				dp[g] = max(maxPath+1, dp[g])
			}
		}
		dpPrev = dp
		dp = make([]int, n+1)
	}
	maxVal := 0
	maxValIndex := 0
	for i, MIV := range dpPrev {
		if MIV > maxVal {
			maxVal = MIV
			maxValIndex = i
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(maxVal))
	writer.WriteByte('\n')
	indexPath := make([]int, 0, maxVal)
	indexPath = append(indexPath, maxValIndex)
	maxVal--
	for i := maxValIndex; i > 0; i-- {
		if maxVal == 0 {
			break
		}
		if dpPrev[i] == maxVal {
			indexPath = append(indexPath, i)
			maxVal--
		}
	}
	indexPath = reverse(indexPath)
	builder := strings.Builder{}
	for _, v := range indexPath {
		builder.WriteString(strconv.Itoa(v) + " ")
	}
	writer.WriteString(builder.String())
	writer.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(v []int) []int {
	for i, g := 0, len(v)-1; i < g; i, g = i+1, g-1 {
		v[i], v[g] = v[g], v[i]
	}
	return v
}

func main() {
	Journey(os.Stdin, os.Stdout)
}
