package crib

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
)

type Prefixes []string

func (p Prefixes) Len() int {
	return len(p)
}

func (p Prefixes) Less(i, j int) bool {
	return len(p[i]) > len(p[j])
}

func (p Prefixes) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// БОР И ПРОХОДИМСЯ ПО ВСЕМ БУКВАМ, В БОРЕ СТАВИМ ХУЙНЮ КАКАУЮ-НИБУДЬ ДЛЯ ОПРЕДЕЛЕНИЯ ВСЕ ЛИ СЛОВА БЫЛИ ЗАЮЗАННЫ

func Crib(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	s, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	nString, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(nString[:len(nString)-1])
	prefixes := make(Prefixes, n)
	for i := 0; i < n; i++ {
		p, _ := reader.ReadString('\n')
		prefixes[i] = p[:len(p)-1]
	}
	sort.Sort(prefixes)
	words := make([]bool, len(s))
	for _, prefix := range prefixes {
		sentinelPrefixLen := len(prefix) + 1
		sentinel := prefix + "#" + s
		dp := make([]int, len(sentinel))
		for i := sentinelPrefixLen; i < len(sentinel); i++ {
			k := dp[i-1]
			for k > 0 && sentinel[k] != sentinel[i] {
				k = dp[k-1]
			}
			if sentinel[i] == sentinel[k] {
				dp[i] = k + 1
				if dp[i] == len(prefix) {
					start := i - len(prefix)*2
					end := i - len(prefix)
					if !words[start] && !words[end-1] {
						for g := start; g < end; g++ {
							words[g] = true
						}
					}
				}
			} else {
				dp[i] = k
			}
		}
	}
	for _, word := range words {
		if !word {
			io.WriteString(w, "NO")
			return
		}
	}
	io.WriteString(w, "YES")
}

func main() {
	Crib(os.Stdin, os.Stdout)
}
