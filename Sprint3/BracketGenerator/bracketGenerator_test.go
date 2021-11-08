package bracketGenerator_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	bracketGenerator "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/BracketGenerator"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBracketGenerator(t *testing.T) {
	tests := []Test{{strings.NewReader("3"), `((()))
(()())
(())()
()(())
()()()
`},
		{strings.NewReader("2"), `(())
()()
`}, {strings.NewReader("4"), `(((())))
((()()))
((())())
((()))()
(()(()))
(()()())
(()())()
(())(())
(())()()
()((()))
()(()())
()(())()
()()(())
()()()()
`}, {strings.NewReader("5"), `((((()))))
(((()())))
(((())()))
(((()))())
(((())))()
((()(())))
((()()()))
((()())())
((()()))()
((())(()))
((())()())
((())())()
((()))(())
((()))()()
(()((())))
(()(()()))
(()(())())
(()(()))()
(()()(()))
(()()()())
(()()())()
(()())(())
(()())()()
(())((()))
(())(()())
(())(())()
(())()(())
(())()()()
()(((())))
()((()()))
()((())())
()((()))()
()(()(()))
()(()()())
()(()())()
()(())(())
()(())()()
()()((()))
()()(()())
()()(())()
()()()(())
()()()()()
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		bracketGenerator.BracketGenerator(v.inputData, buf)

		if buf.String() != v.outputData {
			line := 1
			testS := buf.String()
			for i := 0; i < len(v.outputData); i++ {
				if v.outputData[i] != testS[i] {
					break
				}
				if v.outputData[i] == 10 {
					line++
				}
			}
			fmt.Println(len(testS), len(v.outputData))
			t.Errorf("\nexpected:\n%s\ngot:\n%s\nLINE = %d", v.outputData, buf.String(), line)
		}
	}
}
