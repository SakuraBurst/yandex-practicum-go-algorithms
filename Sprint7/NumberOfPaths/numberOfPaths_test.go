package numberOfPaths_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	numberOfPaths "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/NumberOfPaths"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestNumberOfPaths(t *testing.T) {
	tests := []Test{{strings.NewReader(`3 3
1 2
1 2
2 3
1 3
`), "2"}, {strings.NewReader(`5 3
1 2
3 4
4 5
1 5
`), "0"}, {strings.NewReader(`3 3
1 2
2 3
1 3
1 1
`), "1"}, {strings.NewReader(`10 20
5 4
2 10
2 10
2 3
4 1
5 1
10 6
4 1
1 6
2 10
4 6
9 6
9 10
2 7
4 1
9 1
7 3
2 9
10 4
2 3
9 1`), "4"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			numberOfPaths.NumberOfPaths(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
