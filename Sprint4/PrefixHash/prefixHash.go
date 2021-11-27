package prefixHash

import (
	"bufio"
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

	degreeOfA := nDegreesByModule(a, m, len(s))
	pref := prefixHashesOfString(s, a, m)
	scaner := bufio.NewScanner(reader)

	scaner.Scan()
	t, _ := strconv.Atoi(scaner.Text())
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

func nDegreesByModule(val, mod, n int) []int {
	degrees := make([]int, n+1)
	degrees[0] = 1
	for i := 1; i <= n; i++ {
		res := degrees[i-1] * val
		degrees[i] = res % mod
		// подумоть
		// if val > 0 {

		// } else {
		// 	degrees[i] = ((res % mod) + mod) % mod
		// }
	}
	return degrees
}

func prefixHashesOfString(s string, a, m int) []int {
	pref := make([]int, len(s)+1)
	pref[0] = 0
	for i := 1; i <= len(s); i++ {
		pref[i] = (pref[i-1]*a + int(s[i-1])) % m
	}
	return pref
}

func main() {
	PrefixHash(os.Stdin, os.Stdout)
}
