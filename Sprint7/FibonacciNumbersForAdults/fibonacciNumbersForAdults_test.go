package fibonacciNumbersForAdults_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	fibonacciNumbersForAdults "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/FibonacciNumbersForAdults"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFibonacciNumbersForAdults(t *testing.T) {
	tests := []Test{{strings.NewReader("5\n"), "8"}, {strings.NewReader("2\n"), "2"}, {strings.NewReader("10\n"), "89"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			fibonacciNumbersForAdults.FibonacciNumbersForAdults(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
