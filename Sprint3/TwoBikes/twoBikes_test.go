package twoBikes_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	twoBikes "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/TwoBikes"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestTwoBikes(t *testing.T) {
	tests := []Test{{strings.NewReader(`6
1 2 4 4 6 8
3`), "3 5"}, {strings.NewReader(`6
1 2 4 4 4 4
3`), "3 -1"}, {strings.NewReader(`6
1 2 4 4 4 4
10`), "-1 -1"}, {strings.NewReader(`6
1 1 4 4 4 4
1`), "1 3"}, {strings.NewReader(`6
1 1 4 4 4 4
4`), "3 -1"}, {strings.NewReader(`8
1 11 15 18 21 30 33 39
1`), "1 2"}, {strings.NewReader(`17
12 19 29 31 33 37 52 62 66 68 71 75 75 88 90 93 98
28`), "3 8"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		twoBikes.TwoBikes(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
