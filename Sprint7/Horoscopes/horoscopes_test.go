package horoscopes_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	horoscopes "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Horoscopes"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestHoroscopes(t *testing.T) {
	tests := []Test{{strings.NewReader(`5
4 9 2 4 6
7
9 4 0 0 2 8 4
`), `3
1 3 4
2 5 7`}, {strings.NewReader(`4
1 1 1 1
2
2 2
`), `0`}, {strings.NewReader(`8
1 2 1 9 1 2 1 9
5
9 9 1 9 9`), `3
4 5 8
1 3 4`}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			horoscopes.Horoscopes(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
