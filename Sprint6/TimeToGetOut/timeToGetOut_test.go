package timeToGetOut_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	timeToGetOut "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/TimeToGetOut"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestTimeToGetOut(t *testing.T) {
	tests := []Test{{strings.NewReader(`6 8
2 6
1 6
3 1
2 5
4 3
3 2
1 2
1 4
`), `0 11
1 6
8 9
7 10
2 3
4 5
`}, {strings.NewReader(`3 2
1 2
2 3
`), `0 5
1 4
2 3
`}, {strings.NewReader(`1 0
`), `0 1
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		timeToGetOut.TimeToGetOut(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
