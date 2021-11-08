package bigInt

import (
	"io"
)

func BigInt(r io.Reader, w io.Writer) {}

// package bracketGenerator

// import (
// 	"bufio"
// 	"io"
// 	"sort"
// 	"strconv"
// )

// type BracketSequensesContainer struct {
// 	bracketSequenses []string
// 	howMuchBrackets  int
// }

// func BracketGenerator(r io.Reader, w io.Writer) {
// 	stdReader := bufio.NewReader(r)
// 	nS, _ := stdReader.ReadString('\n')
// 	n, _ := strconv.Atoi(nS)
// 	bufferWriter := bufio.NewWriter(w)
// 	if n == 1 {
// 		bufferWriter.WriteString("()")
// 		bufferWriter.Flush()
// 		return
// 	}
// 	bS := BracketSequensesContainer{make([]string, 0, n*2), n}
// 	recursionBracketAppend(&bS)
// 	sort.Strings(bS.bracketSequenses)

// 	for _, s := range bS.bracketSequenses {
// 		bufferWriter.WriteString(s)
// 		bufferWriter.WriteByte('\n')
// 	}
// 	bufferWriter.Flush()
// }

// func recursionBracketAppend(dc *BracketSequensesContainer) {
// 	if dc.howMuchBrackets == 2 {
// 		dc.bracketSequenses = append(dc.bracketSequenses, "(())", "()()")
// 	} else {
// 		dc.howMuchBrackets--
// 		recursionBracketAppend(dc)
// 		newDecks := make([]string, 0, len(dc.bracketSequenses)*3)
// 		for _, d := range dc.bracketSequenses {
// 			newDecks = append(newDecks, (d + "()"), ("(" + d + ")"))
// 		}
// 		reversedDecks := []string{}
// 		for _, nd := range newDecks {
// 			reverseNds := reverseWithBracketReverse([]byte(nd))
// 			if nd != reverseNds {
// 				reversedDecks = append(reversedDecks, reverseNds)
// 			}
// 		}
// 		dc.bracketSequenses = append(newDecks, reversedDecks...)
// 	}
// }

// func reverseWithBracketReverse(s []byte) string {
// 	for i, g := 0, len(s)-1; i < g; i, g = i+1, g-1 {
// 		s[i], s[g] = reverseBracket(s[g]), reverseBracket(s[i])
// 	}
// 	return string(s)
// }

// func reverseBracket(b byte) byte {
// 	if b == '(' {
// 		return ')'
// 	}
// 	return '('
// }
