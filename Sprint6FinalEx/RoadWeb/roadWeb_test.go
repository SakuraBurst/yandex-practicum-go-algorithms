package roadWeb_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	roadWeb "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6FinalEx/RoadWeb"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestRoadWeb(t *testing.T) {
	tests := []Test{{strings.NewReader(`4 4
1 2 5
1 3 6
2 4 8
3 4 3
`), "19"}, {strings.NewReader(`3 3
1 2 1
1 2 2
2 3 1
`), "3"}, {strings.NewReader(`2 0
`), "Oops! I did it again"}, {strings.NewReader(`10 20
9 10 4
2 2 4
4 2 8
10 5 3
1 10 6
7 4 2
10 10 6
3 7 4
8 9 4
8 10 7
6 10 10
2 8 8
3 8 1
3 10 3
9 5 8
10 10 2
1 8 1
10 1 5
3 6 10
9 10 8
`), "69"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			roadWeb.RoadWeb(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}
