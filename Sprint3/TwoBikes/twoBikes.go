package twoBikes

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func TwoBikes(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	vStr, _ := reader.ReadString('\n')
	v := strings.Fields(strings.TrimSpace(vStr))
	costStr, _ := reader.ReadString('\n')
	cost, _ := strconv.Atoi(strings.TrimSpace(costStr))
	forVasya := binarySearchRec(v, cost, 0, n-1, -1)
	guess, _ := strconv.Atoi(v[forVasya])
	if guess < cost {
		writer.WriteString("-1 -1")
	} else {
		writer.WriteString(strconv.Itoa(forVasya+1) + " ")
		forSister := binarySearchRec(v, cost*2, 0, len(v), -1)
		if forSister == -1 {
			writer.WriteString("-1")
		} else {
			guess, _ := strconv.Atoi(v[forSister])
			if guess >= cost*2 {
				writer.WriteString(strconv.Itoa(forSister + 1))
			} else {
				writer.WriteString("-1")
			}
		}
	}

	writer.Flush()

}

func binarySearchRec(accumulation []string, searchFor int, start, end, lastmid int) int {
	fmt.Println(start, end, lastmid)
	if start >= end {

		currentCheck, _ := strconv.Atoi(accumulation[lastmid])
		if currentCheck > searchFor {
			return lastmid
		} else {
			for lastmid < len(accumulation)-1 {
				nextCheck, _ := strconv.Atoi(accumulation[lastmid])
				if nextCheck < searchFor {
					lastmid++
				} else {
					break
				}
			}
			return lastmid
		}

	}
	middle := (start + end) / 2
	currentCheck, _ := strconv.Atoi(accumulation[middle])
	if currentCheck == searchFor {
		for middle > 0 {
			nextCheck, _ := strconv.Atoi(accumulation[middle-1])
			if nextCheck == searchFor {
				middle--
			} else {
				break
			}
		}
		return middle
	}
	if currentCheck > searchFor {
		return binarySearchRec(accumulation, searchFor, start, middle, middle)
	} else {
		return binarySearchRec(accumulation, searchFor, middle+1, end, middle)
	}
}
