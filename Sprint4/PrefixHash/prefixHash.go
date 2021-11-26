package prefixHash

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PrefixHash(r io.Reader, w io.Writer) {
	// rand.Seed(time.Now().Unix())
	reader := bufio.NewReader(r)
	aString, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(strings.TrimSpace(aString))
	mString, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mString))
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	t, _ := strconv.Atoi(scaner.Text())
	// hashOfIndexRanges := map[[2]int]int{}
	// hashOfIndexRanges[[2]int{0, len(s) - 1}] = hash([]byte(s), a, m)
	degreeOfA := make([]int, 0, t)
	degreeOfA = append(degreeOfA, 1)
	lenS := len(s)
	for i := 1; i <= lenS; i++ {
		degree := degreeOfA[i-1] * a
		if degree < 0 {
			degreeOfA = append(degreeOfA, (degree%m+m)%m)
		} else {
			degreeOfA = append(degreeOfA, degree%m)
		}

	}
	pref := make([]int, lenS+1)
	pref[0] = 0
	for i := 1; i < lenS+1; i++ {
		pref[i] = (int(s[i-1]) + pref[i-1]*a) % m
	}
	fmt.Println(a, m, s, degreeOfA, pref)
	writer := bufio.NewWriter(w)
	for i := 0; i < t; i++ {
		scaner.Scan()
		stringIndexOfString := strings.Fields(scaner.Text())
		aI, _ := strconv.Atoi(stringIndexOfString[0])
		bI, _ := strconv.Atoi(stringIndexOfString[1])
		if aI-1 == 0 {
			writer.WriteString(strconv.Itoa(pref[bI]))
			writer.WriteByte('\n')
		} else {
			v := (pref[bI] - pref[aI-1]*degreeOfA[bI-aI+1])
			fmt.Println(aI, bI, degreeOfA[bI-aI])
			if v < 0 {
				v %= m
				v += m
				v %= m
			}
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte('\n')
		}

	}
	writer.Flush()

}

// func hash(s string, a, m int) int {
// 	// fmt.Println(s)
// 	if len(s) == 1 {
// 		return int(s[0]) % m
// 	}
// 	result := ((int(s[0]) * a) + int(s[1])) % m
// 	for i := 2; i < len(s); i++ {
// 		result = ((result * a) + int(s[i])) % m
// 	}
// 	return result
// }

func main() {
	PrefixHash(os.Stdin, os.Stdout)
}
