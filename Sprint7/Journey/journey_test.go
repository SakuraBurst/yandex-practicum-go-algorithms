package journey_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	journey "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Journey"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestJourney(t *testing.T) {
	tests := []Test{{strings.NewReader(`5
4 2 9 1 13
`), `3
2 3 5 `}, {strings.NewReader(`6
1 2 4 8 16 32
`), `6
1 2 3 4 5 6 `}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			journey.Journey(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
