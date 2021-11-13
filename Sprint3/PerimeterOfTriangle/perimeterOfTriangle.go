package perimeterOfTriangle

import (
	"bufio"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func PerimeterOfTriangle(r io.Reader, w io.Writer) {
	rand.Seed(time.Now().Unix())
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	// n, _ := strconv.Atoi(scaner.Text())
	scaner.Scan()
	data := strings.Fields(scaner.Text())
	intData := make([]int, len(data))
	for i, v := range data {
		n, _ := strconv.Atoi(v)
		intData[i] = n
	}
	intData = quickSort(intData)
	writer := bufio.NewWriter(w)
	for i := 0; i < len(intData); i++ {
		if intData[i] < (intData[i+1] + intData[i+2]) {
			writer.WriteString(strconv.Itoa(intData[i] + intData[i+1] + intData[i+2]))
			writer.Flush()
			return
		}
	}
}

func quickSort(d []int) []int {
	if len(d) < 2 {
		return d
	} else {
		p := pivot(d)
		left, mid, center := partition(d, p)
		result := make([]int, 0, len(d))
		result = append(result, quickSort(left)...)
		result = append(result, mid...)
		result = append(result, quickSort(center)...)
		return result
	}
}

func pivot(v []int) int {
	l := len(v)
	return median(v[rand.Intn(l)], v[rand.Intn(l)], v[rand.Intn(l)])
}

func median(a, b, c int) int {
	if a < b {
		if a > c {
			return a
		}
		if c > b {
			return b
		}
		return c
		// else a > b
	} else {
		if b > c {
			return b
		}
		if a > c {
			return c
		}
		return a
	}
}

func partition(a []int, mid int) ([]int, []int, []int) {
	left := make([]int, 0, len(a)/2)
	right := make([]int, 0, len(a)/2)
	center := make([]int, 0, 1)
	for _, v := range a {
		if v == mid {
			center = append(center, v)
		} else if v > mid {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, center, right
}
